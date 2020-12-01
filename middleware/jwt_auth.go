package middleware

import (
	"call-up/model"
	"call-up/serializer"
	"call-up/service"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

const identityKey = "UID"

var jwtSecret = os.Getenv("JWT_SECRET")

func GinJWTMiddlewareInit() (authMiddleware *jwt.GinJWTMiddleware, err error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "call-up jwt",
		Key:         []byte(jwtSecret),
		Timeout:     3 * time.Hour,
		MaxRefresh:  3 * time.Hour,
		IdentityKey: identityKey,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			//handles the login logic. On success LoginResponse is called, on failure Unauthorized is called
			var serv service.UserLogin
			if err := c.ShouldBind(&serv); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if user, ok := serv.Login(); ok {
				c.Set("user", user)
				return user, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				//maps the claims in the JWT
				return jwt.MapClaims{
					identityKey: strconv.Itoa(int(v.ID)),
				}
			}
			return jwt.MapClaims{}
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			user, _ := c.Get("user")
			c.JSON(code, serializer.BuildUserLoginResponse(user.(*model.User).ID, token, expire))
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			//extracts identity from claims
			//Set the identity
			UID, err := strconv.Atoi(claims[identityKey].(string))
			if err != nil {
				return nil
			}
			user, err := model.GetUser(uint(UID))
			if err != nil {
				return nil
			}
			c.Set("user", &user)
			return &user
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*model.User); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, serializer.Err(serializer.CodeCheckLogin, message, nil))
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			user, _ := c.Get("user")
			c.JSON(code, serializer.BuildUserLoginResponse(user.(*model.User).ID, token, expire))
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}
