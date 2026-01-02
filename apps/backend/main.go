package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
	"os"
	"time"
)

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		items, err := GetItems()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":      "ok",
				"message":   "OK",
				"timestamp": time.Now(),
				"payload":   items,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":      "internal",
				"message":   err.Error(),
				"timestamp": time.Now(),
				"payload":   nil,
			})
		}
	})
	_ = r.Run()
}

func GetItems() ([]Item, error) {
	url := os.Getenv("POSTGRES")
	ctx := context.Background()
	db, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}
	defer db.Close(ctx)
	if err = Migrate(db); err != nil {
		return nil, err
	}
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

func Migrate(db *pgx.Conn) error {
	ctx := context.Background()
	var exists bool
	row := db.QueryRow(ctx, "select exists(select from information_schema.tables where table_schema='public' and table_name='items')")
	if err := row.Scan(&exists); err != nil {
		return err
	}
	if exists {
		return nil
	}
	_, err := db.Exec(ctx, "create table items (id serial primary key, name varchar(255))")
	if err != nil {
		return err
	}
	_, err = db.Exec(ctx, "insert into items (name) values ('A'), ('B'), ('C')")
	if err != nil {
		return err
	}
	return nil
}
