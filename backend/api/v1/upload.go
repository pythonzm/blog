package v1

import (
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// t的选项是avatar或image
func UploadImageAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	t := c.PostForm("t")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(40000, nil, err))
		return
	}

	imageInfo, err := service.UploadImageAvatarService(file, t)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40029, nil, err))
		return
	}
	dst := imageInfo.ImageRelPath
	if dst == "" {
		dst = imageInfo.AvatarRelPath
	}
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40030, nil, err))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, imageInfo, nil))
	return
}
