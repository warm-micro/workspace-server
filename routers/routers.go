package routers

import (
	"wm/workspace/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(router *gin.Engine) {
	router.GET("/ping", controllers.Pong)
	AddWorkspaceRouter(router.Group(""))
}
