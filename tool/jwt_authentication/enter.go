package jwt_authentication

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type CustomClaims struct {
	UserID int32
	jwt.RegisteredClaims
}

// GenerateJWT 签发
func GenerateJWT(userID int32, AccessSecret []byte, AccessExpire time.Duration) (string, error) {
	// 设置Claims
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessExpire)),
			Issuer:    "example.com",
		},
	}

	// 创建Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名Token
	tokenString, err := token.SignedString(AccessSecret)
	return tokenString, err
}

// ParseJWT 解析并验证JWT
func ParseJWT(tokenStr string, AccessSecret []byte) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 确保使用了正确的签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return AccessSecret, nil
	})

	if err != nil {
		// 检查错误类型，提供更具体的错误信息
		if errVal, ok := err.(*jwt.ValidationError); ok {
			switch errVal.Errors {
			case jwt.ValidationErrorExpired:
				return nil, fmt.Errorf("JWT token has expired")
			case jwt.ValidationErrorSignatureInvalid:
				return nil, fmt.Errorf("JWT signature does not match")
			case jwt.ValidationErrorClaimsInvalid:
				return nil, fmt.Errorf("invalid claims")
			case jwt.ValidationErrorMalformed:
				return nil, fmt.Errorf("JWT token is malformed")
			default:
				return nil, fmt.Errorf("unknown error with JWT token")
			}
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid JWT token")
	}
}
