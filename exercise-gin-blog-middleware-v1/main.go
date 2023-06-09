package main

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"
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

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.AbortWithStatus(401)
			return
		}
		split := strings.Split(auth, " ")
		if len(split) != 2 {
			c.AbortWithStatus(401)
			return
		}
		if split[0] != "Basic" {
			c.AbortWithStatus(401)
			return
		}
		decoded, err := base64.StdEncoding.DecodeString(split[1])
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		split = strings.Split(string(decoded), ":")
		if len(split) != 2 {
			c.AbortWithStatus(401)
			return
		}
		username := split[0]
		password := split[1]
		for _, user := range users {
			if user.Username == username && user.Password == password {
				c.Next()
				return
			}
		}
		c.AbortWithStatus(401)
		// fmt.Printf("Authorization: %s\n", auth)
		// fmt.Println(auth)

	} // TODO: replace this
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//Set up authentication middleware here
	rt := r.Group("", authMiddleware())
	{
		rt.GET("/posts", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Query("id"))
			if err != nil && c.Query("id") != "" {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
				return
			}
			if id != 0 {
				for _, post := range Posts {
					if post.ID == id {
						c.JSON(http.StatusOK, gin.H{
							"post": post,
						})
						return
					}
				}
				c.JSON(http.StatusNotFound, gin.H{"error": "Postingan tidak ditemukan"})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"posts": Posts,
			})

		})

		rt.POST("/posts", func(c *gin.Context) {
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
		})
	}
	// r.Use(authMiddleware())

	return r
}

func main() {
	r := SetupRouter()

	r.Run(":8080")
}
