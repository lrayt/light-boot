package http_manager

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lrayt/light-boot/convention"
	"github.com/lrayt/light-boot/core"
	"github.com/lrayt/light-boot/core/log_provider"
)

func ToCTX(c *gin.Context) context.Context {
	var traceId = log_provider.NewTraceId()
	if id, exist := c.Get(log_provider.TraceId); exist {
		if tid, ok := id.(string); ok && len(tid) > 0 {
			traceId = tid
		}
	}
	// uid
	var ctx = context.WithValue(context.Background(), log_provider.TraceId, traceId)
	if uid, exist := c.Get(log_provider.UserId); exist {
		ctx = context.WithValue(ctx, log_provider.UserId, uid)
	}

	// isAdmin
	if isAdmin, exist := c.Get(log_provider.IsAdmin); exist {
		ctx = context.WithValue(ctx, log_provider.IsAdmin, isAdmin)
	} else {
		ctx = context.WithValue(ctx, log_provider.IsAdmin, false)
	}

	// client
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

		c.Next()
	}
}

func GetHttpConf() (*convention.HttpConf, error) {
	var httpConf = new(convention.HttpConf)
	var err = core.GConfigs().PackConf("http", httpConf)
	return httpConf, err
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
