package main

import (
    "github.com/gin-gonic/gin"
    "github.com/rhzx3519/auth-server/jwt"
    "net/http"
)

type Login struct {
    Email    string `json: "email" binding:"required"`
    Password string `json: "password" binding:"required"`
}

func auth(c *gin.Context) {
    var json Login
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    if json.Email != "admin@gmail.com" || json.Password != "123456" {
        c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
        return
    }

    tokenString, err := jwt.Sign(json.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    }
    c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func main() {
    r := gin.Default()
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

        v1.POST("/auth", auth)
    }

    r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
