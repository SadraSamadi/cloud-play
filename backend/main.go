package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/greet", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "World"
		}
		message := fmt.Sprintf("Hello, %s!", name)
		c.JSON(http.StatusOK, gin.H{"message": message})
	})
	addr := fmt.Sprintf("%s:%s", "0.0.0.0", "8080")
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
