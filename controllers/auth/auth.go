package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rhzx3519/auth-server/domain"
	userDB "github.com/rhzx3519/auth-server/persistance/user"
	"github.com/rhzx3519/auth-server/utils/jwt"
	"github.com/rhzx3519/auth-server/utils/salt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const TOKEN_KEY = "Authorization"

func Verify(c *gin.Context) {
	fmt.Printf("X-Original-URI: %v, X-Original-METHOD: %v\n",
		c.GetHeader("X-Original-URI"), c.GetHeader("X-Original-METHOD"))

	var err error
	tokenString := c.GetHeader(TOKEN_KEY)
	if tokenString == "" {
		tokenString, err = c.Cookie(TOKEN_KEY)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "unauthorized",
				"message": "Missing authorization header",
			})
			c.Abort()
			return
		}
	}

	tokenString = tokenString[len("Bearer "):]
	claims, err := jwt.Verify(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "unauthorized",
			"message": "Invalid token",
		})
		c.Abort()
		return
	}
	if _, ok := claims["no"]; !ok {
		fmt.Println("error, cannot find necessary information from  token.")
		c.Abort()
		return
	}
	// Set user info in the request's header
	var claimsJson []byte
	if claimsJson, err = json.Marshal(claims); err != nil {
		fmt.Println("failed to marshal claims", err)
		c.Abort()
		return
	}
	c.Header("X-Forwarded-User", string(claimsJson))
}

type LoginData struct {
	Email    string `json: "email" binding:"required"`
	Password string `json: "password" binding:"required"`
}

type LoginResponse struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
}

func Login(c *gin.Context) {
	var json LoginData
	var err error
	if err = c.ShouldBindJSON(&json); err != nil {
		log.WithError(err).Error("failed to parse json body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user domain.User
	if user, err = userDB.FindUser(json.Email, salt.MD5(json.Password)); err != nil {
		log.WithError(err).Error("no matched user is found")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "no matched user is found"})
		return
	}

	tokenString, err := jwt.Sign(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Writer.Header().Set("Authorization", tokenString)
	c.SetCookie("Authorization", "Bearer "+tokenString, 3600*24, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "You are logged in",
		"data":    LoginResponse{Email: user.Email, Nickname: user.Nickname, Token: tokenString},
	})
}
