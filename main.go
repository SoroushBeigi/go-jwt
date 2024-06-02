package main

import (
	"github.com/SoroushBeigi/go-jwt/initializers"
	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run()
}
