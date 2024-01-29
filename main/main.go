package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rhzx3519/auth-server/controllers/auth"
	"github.com/rhzx3519/auth-server/persistance/mysql"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("DPHOST"), os.Getenv("DBPASS"))

	mysql.InitDB()
}

// This is used to avoid cors(request different domains) problem from the client
func corsHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")

	// When React calls an API, it first sends an OPTIONS request to detect if the API available
	// So return 204 whenever receive an OPTIONS request to avoid CORS error
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
}

func main() {
	r := gin.Default()
	r.Use(corsHeader)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Auth API.\\nPlease use POST /auth & POST /verify for authentication",
			})
		})

		v1.POST("/login", auth.Login)
		v1.POST("/verify", auth.AuthRequired)

		authorized := v1.Group("/", auth.AuthRequired)
		{
			authorized.GET("/testauth", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "You are authorized!",
				})
			})
		}
	}

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
