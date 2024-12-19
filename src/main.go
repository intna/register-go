package main

import (
	"register/src/config"
	"register/src/middleware"
	"register/src/routes"

	"github.com/gin-gonic/gin"
)

var logger = middleware.Log()

func main() {
	config.LoadEnv()

	//connect to mongodb
	mongoURI := config.GetValue("MONGO_URI", "")
	config.ConnectMongoDB(mongoURI)

	// initialize Gin
	r := gin.Default()

	// register routes
	routes.RegisterRoutes(r)

	// start server
	port := config.GetValue("PORT", ":8080")
	if err := r.Run(port); err != nil {

		logger.Fatalf("Failed to start server: %v", err)
	}
}
