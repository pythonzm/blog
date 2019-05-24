package api

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAuth(c *gin.Context) {
	data := make(map[string]interface{})
	username := c.Query("username")
	password := c.Query("password")

	isExist := models.CheckAuth(username, password)
	if isExist {
		token, e := utils.GenToken(username, password)
		if e != nil {
			utils.DefaultResponse(c, http.StatusInternalServerError, utils.TokenGenError, e, nil)
			return
		}
		data["token"] = token
		utils.DefaultResponse(c, http.StatusOK, utils.Success, e, data)
		return
	}
	utils.DefaultResponse(c, http.StatusUnauthorized, utils.AuthError, nil, nil)
	return
}
