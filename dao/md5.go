package dao

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

type MD5 struct {
}

func (m *MD5) encrypt(v string) string {
	d := []byte(v)
	md := md5.New()
	md.Write(d)
	return hex.EncodeToString(md.Sum(nil))
}

func (m *MD5) randomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Int63()%int64(len(charset))]
	}
	return string(b)
}
