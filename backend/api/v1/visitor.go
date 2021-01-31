package v1

import (
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVisitorCount(c *gin.Context) {
	a := service.Visitor{}
	count, e := a.GetVisitorCount()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40037, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, count, nil))
	return
}

func GetCountByDate(c *gin.Context) {
	a := service.Visitor{}
	res, e := a.GetCountByDate()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40037, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, res, nil))
	return
}

func GetCountByUA(c *gin.Context) {
	a := service.Visitor{}
	res, e := a.GetCountByUA()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40037, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, res, nil))
	return
}
