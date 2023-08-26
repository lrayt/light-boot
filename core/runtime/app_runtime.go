package runtime

import (
	"fmt"
	"github.com/light-boot/core/cfg_provider"
	"github.com/light-boot/core/env"
	"github.com/light-boot/core/log_provider"
)

type AppRuntime struct {
	Env            *env.GlobalEnv
	ConfigProvider cfg_provider.ConfigProvider
	LoggerProvider log_provider.LoggerProvider
}

func (e *AppRuntime) SetConfigProvider(configProvider cfg_provider.ConfigProvider) {
	if configProvider != nil {
		e.ConfigProvider = configProvider
	}
}

func (e *AppRuntime) SetLoggerProvider(logger log_provider.LoggerProvider) {
	if logger != nil {
		e.LoggerProvider = logger
	}
}

func (e AppRuntime) Print() {
	fmt.Printf("AppName:%s\n", e.Env.AppName)
	fmt.Printf("RunEnv:%s\n", e.Env.RunEnv)
	fmt.Printf("Version:%s\n", e.Env.BuildVersion)
	fmt.Printf("WorkDir:%s\n", e.Env.WorkDir)
}
