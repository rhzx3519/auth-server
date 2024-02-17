package salt

import (
    "crypto/md5"
    "fmt"
)

const salt = "au"

func MD5(str string) string {
    data := append([]byte(str), []byte(salt)...) //切片
    has := md5.Sum(data)
    md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
    return md5str
}
