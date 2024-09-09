package secret

import (
	"crypto/md5"
	"fmt"
	"io"
)

// Md5 函数将字符串加密为md5，然后返回
func Md5(str string) string {
	hasher := md5.New()
	io.WriteString(hasher, str)
	hash := hasher.Sum(nil)
	return fmt.Sprintf("%s", hash)
}
