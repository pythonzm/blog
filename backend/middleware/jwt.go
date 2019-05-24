package middleware

import (
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
			_, e := utils.ParseToken(token)

			if e != nil {
				code = utils.TokenCheckError
				err = e
			}
		}

		if code != utils.Success {
			utils.DefaultResponse(c, http.StatusUnauthorized, code, err, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
