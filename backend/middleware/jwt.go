package middleware

import (
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWt() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := utils.Success
		var err error

		token := c.Query("token")
		if token == "" {
			code = utils.TokenCheckError
		} else {
			_, e := service.ParseToken(token)

			if e != nil {
				err = e
			}
		}

		if code != utils.Success {
			c.JSON(http.StatusUnauthorized, utils.GenResponse(code, nil, err))
			c.Abort()
			return
		}
		c.Next()
	}
}
