package main

import (
	"github.com/gin-gonic/gin"
	"go-vue-sandbox-api/models"
)

func setupRouter() *gin.Engine {
	// router定義
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// })
	})
	return r
}

func main() {
	r := setupRouter()

	models.ConnectDB()

	r.Run()
	// r.Run(":8080")
}
