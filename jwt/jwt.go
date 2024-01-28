package jwt

import (
    "fmt"
    "github.com/golang-jwt/jwt/v5"
    "time"
)

var jwtSecretKey = []byte("dksjfl93Dds@#@$!sdasd@!#DSSAD")

func Sign(email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256,
        jwt.MapClaims{
            "email": email,
            "exp":   time.Now().Add(time.Hour * 24).Unix(),
        })
    return token.SignedString(jwtSecretKey)
}

func verify(tokenString string) error {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecretKey, nil
    })
    if err != nil {
        return err
    }
    if !token.Valid {
        return fmt.Errorf("invalid token")
    }
    return nil
}
