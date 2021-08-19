package goutils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// 生成token
// claim 是键值对
func GenerateToken(key string, claim map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	if len(claim) > 0 {
		claims := make(jwt.MapClaims)
		for k, v := range claim {
			claims[k] = v
		}
		token.Claims = claims
	}

	tokenString, err := token.SignedString([]byte(key))
	return tokenString, err
}

// 验证token
func ParseToken(token, key string) (map[string]interface{}, error) {
	token1, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(key), nil
	})

	if claims, ok := token1.Claims.(jwt.MapClaims); ok && token1 != nil && token1.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
