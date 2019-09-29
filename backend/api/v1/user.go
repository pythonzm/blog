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
	c.JSON(http.StatusOK, utils.GenResponse(20000, nil, nil))
}

func GetUserInfo(c *gin.Context) {
	userInfo, e := service.GetUser()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40027, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, userInfo, nil))
	return
}

func GetUserAbout(c *gin.Context) {
	about, e := service.GetAbout()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40027, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, about, nil))
	return
}

func EditUser(c *gin.Context) {

	bytes, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40028, nil, err))
		return
	}

	u := service.User{}
	if e := json.Unmarshal(bytes, &u); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40028, nil, e))
		return
	}

	if u.Password != "" {
		if e := u.ResetPassword(); e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40028, nil, e))
			return
		}
	} else if u.About != "" {
		if e := u.EditAbout(); e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40028, nil, e))
			return
		}
	} else {
		if e := u.EditUser(); e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40028, nil, e))
			return
		}
	}

	c.JSON(http.StatusOK, utils.GenResponse(20000, u, nil))
	return
}
