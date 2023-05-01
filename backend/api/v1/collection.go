package v1

import (
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCollection(c *gin.Context) {
	collection, e := service.Collection{}.GetCollection()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40040, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, collection, nil))
}

func CudCollection(c *gin.Context) {
	collection := &service.Collection{}
	if e := c.ShouldBindJSON(collection); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(40000, nil, e))
		return
	}
	t, e := collection.CudCollection()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40041, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, t, nil))
}
