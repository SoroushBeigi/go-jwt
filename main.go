package main

import (
	"github.com/SoroushBeigi/go-jwt/controllers"
	"github.com/SoroushBeigi/go-jwt/initializers"
	"github.com/SoroushBeigi/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("login", controllers.Login)
	r.GET("validate", middleware.RequireAuth,controllers.Validate)
	r.Run()
}
