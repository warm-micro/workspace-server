package routers

import (
	"wm/workspace/controllers"

	"github.com/gin-gonic/gin"
)

func AddWorkspaceRouter(router *gin.RouterGroup) {
	router.GET("", controllers.ListWorkspace)
	router.POST("/", controllers.CreateWorkspace)
	router.DELETE("/:workspaceId", controllers.DeleteWorkspace)
	router.GET("/exists", controllers.CheckWorkspace)
	router.POST("/inviteCode/", controllers.CreateInviteCode)
	router.POST("/accept", controllers.AcceptInvite)
}
