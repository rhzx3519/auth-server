package jwt

import (
    "gotest.tools/v3/assert"
    "testing"
)

func TestExtractInformation(t *testing.T) {
    var email, no = "eren@gmail.com", "eiwieo493kbbfd"
    tokenString, err := Sign(email, no)
    assert.ErrorIs(t, err, nil)
    claims, err := Verify(tokenString)
    assert.ErrorIs(t, err, nil)
    assert.Equal(t, claims["email"].(string), email)
    assert.Equal(t, claims["no"].(string), no)
}
