package core

import (
	"context"
	"github.com/lrayt/light-boot/core/log_provider"
)

// GAppName 运行服务名
func GAppName() string {
	return appRuntime.Env.AppName
}

func GRunEnv() string {
	return string(appRuntime.Env.RunEnv)
}

// GBuildVersion 构建版本
func GBuildVersion() string {
	return appRuntime.Env.BuildVersion
}

// NewLoggerWithCTX 日志收集器
func NewLoggerWithCTX(module string, args ...context.Context) log_provider.Logger {
	commonFields := map[string]interface{}{
		"run_env": GRunEnv(),
		"version": GBuildVersion(),
		"module":  module,
	}
	if len(args) > 0 {
		commonFields[log_provider.TraceId] = log_provider.GetTraceId(args[0])
		commonFields[log_provider.BizId] = log_provider.GetBizId(args[0])
	}

	return appRuntime.LoggerProvider.NewLogger(commonFields)
}
