package user

import (
	"github.com/rhzx3519/auth-server/persistance/mysql"
	"github.com/rhzx3519/auth-server/utils/salt"
	"gotest.tools/v3/assert"
	"testing"
)

func TestGetUserbyEmailAndPassword(t *testing.T) {
	mysql.InitDB()
	defer func() {
		mysql.PostDB()
	}()

	user, err := FindUser("lou@gmail.com", salt.MD5("123456"))
	assert.NilError(t, err)
	assert.Equal(t, user.Nickname, "lou")
}
