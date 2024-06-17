package encrypt_AES

import (
	"crypto/rand"
	"fmt"
	"io"
	"testing"
)

func generateKey() ([]byte, error) {
	key := make([]byte, 32) // AES-256 需要 32 字节长的密钥
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, err
	}
	return key, nil
}

func TestEncryptAES(t *testing.T) {
	key := "sdjajgkejaiodfackjakjgeiadeklesb"
	plaintext := "18108238484"
	aes, err := AesEncrypt([]byte(plaintext), []byte(key))
	fmt.Println(string(aes), err)
	res, err := AesDecrypt(aes, []byte(key))
	fmt.Println(string(res), err)
}
