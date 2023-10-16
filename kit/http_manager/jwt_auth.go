package http_manager

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lrayt/light-boot/core/log_provider"
	"time"
)

const JWTSigned = "mini-oss:lrayt"

type JWTUserClaims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}

func JWTUser2Token(user *JWTUserClaims) (string, error) {
	claims := JWTUserClaims{
		UserId:   user.UserId,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
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

// JWTAuthMiddleware 鉴权
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token = c.Request.Header.Get("Authorization")
		if len(token) <= 0 {
			c.AbortWithStatus(401)
			return
		}

		if user, err := JWTToken2User(token); err != nil {
			c.AbortWithStatus(401)
			return
		} else {
			c.Set(log_provider.UserId, user.UserId)
			c.Set(log_provider.IsAdmin, user.IsAdmin)
		}
		// 放行
		c.Next()
	}
}
