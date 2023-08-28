package http_manager

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lrayt/light-boot/core/log_provider"
)

func ToCTX(c *gin.Context) context.Context {
	var traceId = log_provider.NewTraceId()
	if id, exist := c.Get(log_provider.TraceId); exist {
		if tid, ok := id.(string); ok && len(tid) > 0 {
			traceId = tid
		}
	}
	return context.WithValue(context.Background(), log_provider.TraceId, traceId)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Signature, X-authorize-uuid")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		var traceId = c.GetHeader(log_provider.RequestId)

		if len(traceId) <= 0 {
			traceId = log_provider.NewTraceId()
		}
		c.Set(log_provider.TraceId, traceId)
		c.Writer.Header().Set(log_provider.RequestId, traceId)
		c.Next()
	}
}
