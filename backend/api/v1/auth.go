package v1

import (
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAuth(c *gin.Context) {
	data := make(map[string]interface{})
	username := c.PostForm("username")
	password := c.PostForm("password")

	a := service.Auth{Username: username, Password: password}
	isExist := a.CheckAuth()
	if isExist {
		token, e := a.GenToken()
		if e != nil {
			c.JSON(http.StatusOK, utils.GenResponse(40004, nil, e))
			return
		}
		data["token"] = token
		c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(40001, nil, nil))
	return
}
