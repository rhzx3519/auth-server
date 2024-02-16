package jwt

import (
    "github.com/rhzx3519/auth-server/domain"
    "gotest.tools/v3/assert"
    "testing"
)

func TestExtractInformation(t *testing.T) {
    var email, no = "eren@gmail.com", "eiwieo493kbbfd"
    tokenString, err := Sign(&domain.User{
        Email:    email,
        Nickname: "eren",
        No:       no,
    })
    assert.ErrorIs(t, err, nil)
    claims, err := Verify(tokenString)
    assert.ErrorIs(t, err, nil)
    assert.Equal(t, claims["email"].(string), email)
    assert.Equal(t, claims["no"].(string), no)
}
