package jwt_authentication

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	AccessSecret := "awejlajwgkawefkdkgca"
	AccessExpire := 3600
	jwt, err := GenerateJWT(1, []byte(AccessSecret), time.Duration(AccessExpire)*time.Second)
	fmt.Println(jwt)
	parseJWT, err := ParseJWT(jwt, []byte(AccessSecret))
	fmt.Println(parseJWT.UserID, err)
}
