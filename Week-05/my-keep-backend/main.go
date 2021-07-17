package main

import (
	"time"

	"my-keep-backend/db"
	"my-keep-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"xorm.io/builder"
)

/*
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	//temp := v.X*v.X + v.Y*v.Y

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
	//(30 * 30) + (40 * 40)
}

func (v *Vertex) Scale(f float64) {
	fmt.Println("X=", v.X)
	fmt.Println("Y=", v.X)
	v.X = v.X * f // X= 3 * 10
	v.Y = v.Y * f // Y = 4 * 10
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.X)
	fmt.Println(v.Y)
	v.Scale(10)
	fmt.Println(v.Abs())
}
*/
var dbClient db.Db

func main() {
	dbClient = db.Db{}
	dbClient.Connect("postgres", "P@ssw0rd", "localhost", 5432, "keep", "postgres")
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		username := "tanto"
		print(username)
		c.JSON(200, gin.H{
			"message": "pong",
			"time":    time.Now(),
		})
	})

	route.POST("/note", func(c *gin.Context) {
		var note models.Note
		var err error
		newId, _ := uuid.NewUUID()
		note.Id = newId.String()
		c.Bind(&note)
		_, err = dbClient.Conn.Insert(&note)
		if err != nil {
			c.JSON(503, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, note)
		}
	})
	route.GET("/note/:id", func(c *gin.Context) {
		id := c.Param("id")
		var note models.Note
		_, err := dbClient.Conn.Where(builder.Eq{"id": id}).Get(&note)
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
