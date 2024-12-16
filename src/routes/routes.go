package routes

import (
	"github.com/gin-gonic/gin"

	"register/src/controller"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/")
	{
		api.POST("/register", controller.RegisterHandler)
		api.POST("/logout", controller.LogoutHandler)
	}
}
