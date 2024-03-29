package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rhzx3519/auth-server/domain"
	"time"
)

var jwtSecretKey = []byte("dksjfl93Dds@#@$!sdasd@!#DSSAD")

const EXPIRED_DURATION = time.Hour * 24

func Sign(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":    user.Email,
			"no":       user.No,
			"nickname": user.Nickname,
			"fullname": user.Firstname + " " + user.Lastname,
			"exp":      time.Now().Add(EXPIRED_DURATION).Unix(),
		})
	return token.SignedString(jwtSecretKey)
}

func Verify(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			return nil, fmt.Errorf("token is expired")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("failed to parse token claims.")
}
