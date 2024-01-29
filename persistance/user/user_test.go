package user

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestGetUserbyEmailAndPassword(t *testing.T) {
	user, err := GetUserbyEmailAndPassword("admin@gmail.com", "123456")
	assert.NilError(t, err)
	assert.Equal(t, user.Nickname, "admin")
}
