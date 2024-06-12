package auth

import (
	"time"
)

const tmp = "0123456789"
const EmailAuthLen = 6

func CreateAuth(len int) string {
	t := time.Now().UnixNano()
	t >>= 5

	code := make([]byte, len)
	for i := 0; i < len; i++ {
		code[i] = tmp[t%10]
		t >>= 1
	}
	return string(code)
}
