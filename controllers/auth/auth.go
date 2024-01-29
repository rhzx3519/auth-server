package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/rhzx3519/auth-server/persistance/user"
	"github.com/rhzx3519/auth-server/utils/jwt"
	"github.com/rhzx3519/auth-server/utils/salt"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "unauthorized",
			"message": "Missing authorization header",
		})
		c.Abort()
		return
	}
	tokenString = tokenString[len("Bearer "):]
	if err := jwt.Verify(tokenString); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "unauthorized",
			"message": "Invalid token",
		})
		c.Abort()
		return
	}
}

type LoginData struct {
	Email    string `json: "email" binding:"required"`
	Password string `json: "password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if _, err := user.GetUserbyEmailAndPassword(json.Email, salt.MD5(json.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	tokenString, err := jwt.Sign(json.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Writer.Header().Set("Authorization", tokenString)
	c.JSON(http.StatusOK, gin.H{
		"message": "You are logged in",
		"token":   tokenString,
	})
}
