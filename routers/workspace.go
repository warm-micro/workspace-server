package routers

import (
	"wm/workspace/controllers"

	"github.com/gin-gonic/gin"
)

func AddWorkspaceRouter(router *gin.RouterGroup) {
	router.GET("", controllers.ListWorkspace)
	router.POST("/", controllers.CreateWorkspace)
}
