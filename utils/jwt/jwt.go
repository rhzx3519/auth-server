package jwt

import (
    "fmt"
    "github.com/golang-jwt/jwt/v5"
    "time"
)

var jwtSecretKey = []byte("dksjfl93Dds@#@$!sdasd@!#DSSAD")

const EXPIRED_DURATION = time.Hour * 24

func Sign(email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256,
        jwt.MapClaims{
            "email": email,
            "exp":   time.Now().Add(EXPIRED_DURATION).Unix(),
        })
    return token.SignedString(jwtSecretKey)
}

func Verify(tokenString string) error {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecretKey, nil
    })
    if err != nil {
        return err
    }
    if !token.Valid {
        return fmt.Errorf("invalid token")
    }
    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        if claims["exp"].(float64) < float64(time.Now().Unix()) {
            return fmt.Errorf("token is expired")
        }
        fmt.Printf("%s is verified.\n", claims["email"].(string))
    }

    return nil
}
