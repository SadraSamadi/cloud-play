package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
	"os"
)

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	//r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		items, err := GetItems()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"items": items})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	_ = r.Run()
}

func GetItems() ([]Item, error) {
	url := os.Getenv("DB_URL")
	ctx := context.Background()
	db, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}
	defer db.Close(ctx)
	rows, err := db.Query(ctx, "select id, name from items")
	if err != nil {
		return nil, err
	}
	var items []Item
	for rows.Next() {
		item := Item{}
		err = rows.Scan(&item.Id, &item.Name)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}
