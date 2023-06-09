package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var Posts = []Post{
	{ID: 1, Title: "Judul Postingan Pertama", Content: "Ini adalah postingan pertama di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Title: "Judul Postingan Kedua", Content: "Ini adalah postingan kedua di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/posts", func(c *gin.Context) {
		fmt.Println(len(Posts))
		c.JSON(http.StatusOK, gin.H{
			"posts": Posts,
		})
	})

	r.GET("/posts/:id", func(c *gin.Context) {

		// if !unicode.IsDigit(rune(c.Param("id")[0])) {
		// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID harus berupa angka"})
		// 	return
		// }
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
			return
		}

		for _, post := range Posts {
			if post.ID == id {
				c.JSON(http.StatusOK, gin.H{
					"post": post,
				})
				return
			}
		}
		// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Postingan tidak ditemukan"})
		c.JSON(http.StatusNotFound, gin.H{"error": "Postingan tidak ditemukan"})
	})

	r.POST("/posts", func(c *gin.Context) {
		var newPost Post
		if err := c.BindJSON(&newPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		newPost.ID = len(Posts) + 1
		Posts = append(Posts, newPost)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Postingan berhasil ditambahkan",
			"post":    newPost,
		})
		// TODO: answer here
	})

	return r
}

func main() {
	r := SetupRouter()

	r.Run(":8080")
}
