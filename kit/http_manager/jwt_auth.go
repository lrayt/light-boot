package http_manager

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const JWTSigned = "mini-oss:lrayt"

type JWTUserClaims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Enable   bool   `json:"enable"`
	jwt.StandardClaims
}

func JWTUser2Token(user *JWTUserClaims) (string, error) {
	claims := JWTUserClaims{
		UserId:   user.UserId,
		Username: user.Username,
		Enable:   true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 120).Unix(),
			Issuer:    "frank",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JWTSigned))
}

func JWTToken2User(token string) (*JWTUserClaims, error) {
	tokenClaims, parseErr := jwt.ParseWithClaims(token, &JWTUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSigned), nil
	})

	if parseErr != nil || tokenClaims == nil {
		return nil, errors.New(fmt.Sprintf("parse token err:%v", parseErr))
	}

	if claims, ok := tokenClaims.Claims.(*JWTUserClaims); !ok || !tokenClaims.Valid {
		return nil, errors.New("token is invalid")
	} else {
		return claims, nil
	}
}
