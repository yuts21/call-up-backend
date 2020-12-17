package service

import (
	"call-up/cache"
	"call-up/model"
	"call-up/serializer"
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CallupRequestHandle 召集令请求处理服务
type CallupRequestHandle struct {
	RequestID   uint  `form:"request_id" json:"request_id" binding:"required"`
	Instruction uint8 `form:"inst" json:"inst" binding:"required,gte=1,lte=2"`
}

// Handle 处理接令请求
func (service *CallupRequestHandle) Handle(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var request model.Request
	if err := model.DB.Where("id = ?", service.RequestID).First(&request).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求查询失败", err)
	}

	var callup model.Callup
	if err := model.DB.Model(&request).Association("Callup").Find(&callup); err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令查询失败", err)
	}

	if callup.SponsorID != user.ID {
		return serializer.Err(serializer.CodeNoRightErr, "无权限", nil)
	}

	if request.Status != model.Unprocessed {
		return serializer.Err(serializer.CodeParamErr, "接令请求不在待处理状态", nil)
	}

	status := callup.Status()
	if status == model.Expired || status == model.Canceled {
		return serializer.Err(serializer.CodeParamErr, "召集令不可用", nil)
	}

	if service.Instruction == 2 {
		if err := model.DB.Model(&request).Update("status", model.Denied).Error; err != nil {
			return serializer.Err(serializer.CodeDBError, "操作失败", err)
		}
		return serializer.Success("操作成功")
	}

	if status == model.Completed {
		return serializer.Err(serializer.CodeParamErr, "召集令人数已满", nil)
	}

	tx := model.DB.Begin()
	if err := tx.Model(&request).Update("status", model.Agreed).Error; err != nil {
		tx.Rollback()
		return serializer.Err(serializer.CodeDBError, "操作失败", err)
	}

	var count int64 = 0
	strCallupID := "callup_" + strconv.FormatUint(uint64(callup.ID), 10)
	if callupKeyExists, err := cache.RedisClient.Exists(strCallupID).Result(); err != nil {
		tx.Rollback()
		return serializer.Err(serializer.CodeCacheError, "缓存访问失败", err)
	} else {
		if callupKeyExists != 0 {
			count, _ = strconv.ParseInt(cache.RedisClient.Get(strCallupID).Val(), 10, 64)
			count++
		} else {
			if err := tx.Model(&model.Request{}).Where("callup_id = ? and status = ?", callup.ID, model.Agreed).Count(&count).Error; err != nil {
				tx.Rollback()
				return serializer.Err(serializer.CodeDBError, "召集令请求查询失败", err)
			}
		}
	}
	if err := cache.RedisClient.Set(strCallupID, strconv.FormatInt(count, 10), 0).Err(); err != nil {
		tx.Rollback()
		return serializer.Err(serializer.CodeCacheError, "操作失败", err)
	}

	if uint(count) >= callup.Capacity {
		date := time.Now()
		date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
		successCallupDetail := model.SuccessCallupDetail{
			CallupID: callup.ID,
			Date:              date,
			SponsorProfit:     3 * uint(count),
			ParticipantProfit: 1 * uint(count),
		}
		if err := tx.Create(&successCallupDetail).Error; err != nil {
			tx.Rollback()
			return serializer.Err(serializer.CodeDBError, "操作失败", err)
		}

		var agencyProfit model.AgencyProfit
		if err := tx.Where("success_date = ? and province = ? and city = ? and type = ?", date, user.Province, user.City, callup.Type).First(&agencyProfit).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				tx.Rollback()
				return serializer.Err(serializer.CodeDBError, "中介收益查询失败", err)
			} else {
				agencyProfit := model.AgencyProfit{
					SuccessDate:  date,
					Province:     user.Province,
					City:         user.City,
					Type:         callup.Type,
					CompletedNum: 1,
					Profit:       successCallupDetail.SponsorProfit + successCallupDetail.ParticipantProfit,
				}
				if err := tx.Create(&agencyProfit).Error; err != nil {
					tx.Rollback()
					return serializer.Err(serializer.CodeDBError, "中介收益创建失败", err)
				}
			}
		} else {
			agencyProfitNew := make(map[string]interface{})
			agencyProfitNew["completed_num"] = agencyProfit.CompletedNum + 1
			agencyProfitNew["profit"] = agencyProfit.Profit + successCallupDetail.SponsorProfit + successCallupDetail.ParticipantProfit
			if err := tx.Model(&agencyProfit).Updates(agencyProfitNew).Error; err != nil {
				tx.Rollback()
				return serializer.Err(serializer.CodeDBError, "中介收益修改失败", err)
			}
		}
	}

	tx.Commit()
	return serializer.Success("操作成功")
}
