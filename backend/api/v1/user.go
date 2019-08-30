package v1

import (
	"backend/service"
	"backend/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	data := make(map[string]interface{})
	user := service.User{}
	if e := c.ShouldBindJSON(&user); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(40000, nil, e))
		return
	}
	isExist := user.CheckAuth()
	if isExist {
		token, e := user.GenToken()
		if e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40004, nil, e))
			return
		}
		data["token"] = token
		c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
		return
	}
	c.JSON(http.StatusUnauthorized, utils.GenResponse(40001, nil, nil))
	return
}

func Logout(c *gin.Context) {
	service.Logout()
	c.JSON(http.StatusOK, utils.GenResponse(20000, nil, nil))
}

func GetUserInfo(c *gin.Context) {
	token := c.Query("token")
	userInfo, e := service.GetUserInfo(token)
	if e != nil {
		c.JSON(http.StatusUnauthorized, utils.GenResponse(40027, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, userInfo, nil))
	return
}

func EditUser(c *gin.Context) {
	token := c.Query("token")
	userInfo, e := service.GetUserInfo(token)
	if e != nil {
		c.JSON(http.StatusUnauthorized, utils.GenResponse(40027, nil, e))
		return
	}

	bytes, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40028, nil, e))
		return
	}

	u := service.User{}
	if e := json.Unmarshal(bytes, &u); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40028, nil, e))
		return
	}
	u.Username = userInfo.Username
	if u.Password == "" {
		if e := u.EditUser(); e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40028, nil, e))
			return
		}
	} else {
		if e := u.ResetPassword(); e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40028, nil, e))
			return
		}
	}

	c.JSON(http.StatusOK, utils.GenResponse(20000, u, nil))
	return
}
