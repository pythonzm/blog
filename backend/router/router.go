package router

import (
	"backend/api/v1"
	"backend/middleware"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
		return
	})
	r.Static(utils.AppInfo.StaticBasePath, utils.AppInfo.UploadBasePath)
	r.Use(middleware.CustomLogger(), middleware.CorsMiddleware())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")

	{
		apiv1.POST("/user/login", v1.Login)
		apiv1.POST("/user/logout", v1.Logout)
		// GET方法没有使用JWT认证
		apiv1.GET("/user/info", v1.GetUserInfo)
		apiv1.GET("/tags", v1.GetAllTags)
		apiv1.GET("/categories", v1.GetAllCategories)
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)

		apiv1.Use(middleware.JWt())
		apiv1.PATCH("/user/edit", v1.EditUser)

		apiv1.POST("/tags", v1.CreateTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.POST("/categories", v1.CreateCategory)
		apiv1.PUT("/categories/:id", v1.EditCategory)
		apiv1.DELETE("/categories/:id", v1.DeleteCategory)


		apiv1.POST("/articles", v1.CreateArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)

		apiv1.POST("/upload", v1.UploadImageAvatar)
	}

	return r
}
