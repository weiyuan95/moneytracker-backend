package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"moneytracker-backend/handlers"
	"net/http"
)

// trigger ci
func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/rate", handlers.Rate)
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	return router
}

func main() {
	router := SetupRouter()

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Unable to start server: ", err)
	}
}
