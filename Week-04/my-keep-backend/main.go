package main

import (
	"time"

	"my-keep-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func main() {
	route := gin.Default()
	route.Any("/ping", func(c *gin.Context) {
		username := "tanto"
		print(username)
		c.JSON(200, gin.H{
			"message": "pong",
			"time":    time.Now(),
		})
	})

	route.POST("/note", func(c *gin.Context) {
		var note models.Note
		newId, _ := uuid.NewUUID()
		note.Id = newId.String()
		c.Bind(&note)
		engine, err := xorm.NewEngine("postgres", "host=localhost port=5432 user=postgres password=P@ssw0rd dbname=keep sslmode=disable")
		_, err = engine.Insert(&note)
		if err != nil {
			c.JSON(503, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, note)
		}
	})

	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
