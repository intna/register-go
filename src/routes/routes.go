package routes

import (
	"github.com/gin-gonic/gin"

	"register/src/controller"
	"register/src/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/")
	{
		api.POST("/register", middleware.CatchErrors(controller.RegisterHandler))
		api.POST("/logout", middleware.CatchErrors(controller.LogoutHandler))
	}
}
