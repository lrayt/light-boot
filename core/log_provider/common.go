package log_provider

import (
	"context"
	"github.com/lrayt/light-boot/pkg/uuid"
)

const (
	RequestId = "Request-Id"
	TraceId   = "trace-id"
	BizId     = "biz-id"
	UserId    = "user-id"
	UserName  = "username"
	IsAdmin   = "is_admin"
	ClientId  = "Client-Id"
)

func NewTraceId() string {
	return uuid.GenUUIDWithPrefix("trace")
}

func GetRequestId(ctx context.Context) string {
	if id, ok := ctx.Value(RequestId).(string); ok {
		return id
	} else {
		return "no-request-id"
	}
}

func GetTraceId(ctx context.Context) string {
	if id, ok := ctx.Value(TraceId).(string); ok {
		return id
	} else {
		return "no-trace-id"
	}
}

func GetBizId(ctx context.Context) string {
	if id, ok := ctx.Value(BizId).(string); ok {
		return id
	} else {
		return "no-biz-id"
	}
}
