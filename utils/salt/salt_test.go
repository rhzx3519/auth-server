package salt

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestMD5(t *testing.T) {
	s := MD5("123456")
	assert.Equal(t, s, "e10adc3949ba59abbe56e057f20f883e")
}
