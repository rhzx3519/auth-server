package salt

import (
    "gotest.tools/v3/assert"
    "testing"
)

func TestMD5(t *testing.T) {
    s := MD5("123456")
    assert.Equal(t, s, "65e5c44190b8fb0cb85840f3684995e8")
}
