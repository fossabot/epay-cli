package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) []byte {
	m := md5.New()
	m.Write([]byte(str))
	return m.Sum(nil)
}

func MD5String(str string) string {
	return fmt.Sprintf("%x", MD5(str))
}
