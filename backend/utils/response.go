package utils

import (
	"github.com/gin-gonic/gin"
)

func DefaultResponse(c *gin.Context, httpCode, errCode int, err error, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  msg(errCode, err),
		"data": data,
	})
}

func Response(c *gin.Context, httpCode, errCode int, err error, data interface{}, total int) {
	c.JSON(httpCode, gin.H{
		"code":  errCode,
		"msg":   msg(errCode, err),
		"data":  data,
		"total": total,
	})
}

func msg(code int, e error) (r string) {
	if code == Success || e == nil {
		r = ErrorText(code)
		return
	}
	r = ErrorText(code) + ": " + e.Error()
	return
}
