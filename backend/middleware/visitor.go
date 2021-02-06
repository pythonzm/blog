package middleware

import (
	"backend/service"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Visitor() gin.HandlerFunc {
	return func(context *gin.Context) {
		substrs := []string{"token", "login", "logout", "info", "soup", "comment", "visitor", "favicon", "static"}
		if !utils.IfContainStr(context.Request.RequestURI, substrs) {
			visitor := &service.Visitor{Ua: context.Request.UserAgent(), IP: context.ClientIP(), Uri: context.Request.RequestURI, Referer: context.Request.Referer()}
			e := visitor.NewVisitor()
			if e != nil {
				utils.WriteErrorLog(fmt.Sprintf("添加访客失败：%v", e))
			}
		}
		context.Next()
	}
}
