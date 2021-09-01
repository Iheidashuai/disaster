package utils

import (
	"crypto/md5"
	"fmt"
)

func Encrypt(origin string) string {
	data := []byte(origin)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
