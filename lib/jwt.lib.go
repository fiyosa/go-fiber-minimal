package lib

import (
	"errors"
	"fmt"
	"go-fiber-minimal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Jwt JwtMiddleware

type JwtMiddleware struct{}

// "data" => "hash_token"
func (*JwtMiddleware) Create(data string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(time.Second * 60 * 60 * 24).Unix(), // 24 hours
	})
	tokenHash, err := token.SignedString([]byte(config.Env.APP_SECRET))
	if err != nil {
		return "", err
	}
	return tokenHash, nil
}

// "hash_token" => "data"
func (*JwtMiddleware) Verify(token string) (string, error) {
	getToken, _ := jwt.Parse(token, func(getToken *jwt.Token) (any, error) {
		if _, ok := getToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %v", getToken.Header["alg"])
		}
		return []byte(config.Env.APP_SECRET), nil
	})
	claims, ok := getToken.Claims.(jwt.MapClaims)
	if !ok || !getToken.Valid {
		return "", errors.New("Unauthorized")
	}
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return "", errors.New("Unauthorized")
	}
	return claims["data"].(string), nil
}
