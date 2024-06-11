package http_manager

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lrayt/light-boot/convention"
	"github.com/lrayt/light-boot/core"
	"github.com/lrayt/light-boot/core/log_provider"
	"github.com/lrayt/light-boot/pkg/uuid"
)

func ToCTX(c *gin.Context) context.Context {
	var ctx = context.Background()
	ctx = context.WithValue(ctx, log_provider.TraceId, c.GetString(log_provider.TraceId))
	ctx = context.WithValue(ctx, log_provider.UserId, c.GetString(log_provider.UserId))
	ctx = context.WithValue(ctx, log_provider.UserName, c.GetString(log_provider.UserName))
	ctx = context.WithValue(ctx, log_provider.IsAdmin, c.GetBool(log_provider.IsAdmin))
	ctx = context.WithValue(ctx, log_provider.ClientId, c.GetHeader(log_provider.ClientId))
	return ctx
}

// CORSMiddleware 跨域设置
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Signature, X-authorize-uuid, Client-Id")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		var traceId = uuid.GenUUID()
		c.Set(log_provider.TraceId, traceId)
		c.Writer.Header().Set(log_provider.RequestId, traceId)
		c.Next()
	}
}

func GetHttpConf() (*convention.HttpConf, error) {
	var httpConf = new(convention.HttpConf)
	var err = core.GConfigs().PackConf("http", httpConf)
	return httpConf, err
}

func CurrUserId(ctx context.Context) string {
	uid := ctx.Value(log_provider.UserId)
	if str, ok := uid.(string); ok && len(str) > 0 {
		return str
	} else {
		return ""
	}
}

func CurrUsername(ctx context.Context) string {
	username := ctx.Value(log_provider.UserName)
	if str, ok := username.(string); ok && len(str) > 0 {
		return str
	} else {
		return ""
	}
}

func CurrClientId(ctx context.Context) string {
	clientId := ctx.Value(log_provider.ClientId)
	if str, ok := clientId.(string); ok && len(str) > 0 {
		return str
	} else {
		return ""
	}
}

func GetUserId(ctx context.Context) (string, error) {
	uid := ctx.Value(log_provider.UserId)
	if str, ok := uid.(string); ok && len(str) > 0 {
		return str, nil
	} else {
		return "", errors.New("找不到UID")
	}
}

func IsAdmin(ctx context.Context) bool {
	isAdmin := ctx.Value(log_provider.IsAdmin)
	is, ok := isAdmin.(bool)
	return ok && is
}

func ClientId(ctx context.Context) string {
	cid := ctx.Value(log_provider.ClientId)
	if clientId, ok := cid.(string); ok {
		return clientId
	} else {
		return "None"
	}
}
