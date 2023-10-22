package jwt

import (
	"FileServerWeb/config"
	"errors"

	"github.com/golang-jwt/jwt/v5"

	"strings"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims  // v5版本新加的方法
}

func GenerateToken(username string) (string, error) {
	var current = time.Now()
	var claim = Claims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(current.Add(3 * time.Hour)), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(current), // 签发时间
			NotBefore: jwt.NewNumericDate(current), // 生效时间
		},
	}

	var t = jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	var s, err = t.SignedString(config.SECRET_KEY)

	return "Bearer " + s, err
}

func ParseToken(s string) (*Claims, error) {
	if s == "" {
		return nil, errors.New("Not login")
	}
	var res = strings.Split(s, " ")

	if res[0] != "Bearer" {
		return nil, errors.New("Wrong format")
	}

	var t, err = jwt.ParseWithClaims(
		res[1],
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.SECRET_KEY), nil
	})

	var claims, ok = t.Claims.(*Claims)
	if ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
