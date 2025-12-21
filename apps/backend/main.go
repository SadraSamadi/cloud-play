package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"items": [...]map[string]any{
				{"id": 1, "title": "Item 1"},
				{"id": 2, "title": "Item 2"},
				{"id": 3, "title": "Item 3"},
			},
		})
	})
	_ = r.Run()
}
