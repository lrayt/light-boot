package core

import (
	"errors"
	"github.com/light-boot/core/cfg_provider"
	"github.com/light-boot/core/env"
	"github.com/light-boot/core/log_provider"
	"github.com/light-boot/core/runtime"
	"github.com/light-boot/pkg/file_utils"
)

var appRuntime *runtime.AppRuntime

// InitApp 初始化应用
func InitApp(appName, workDir string, args ...string) error {
	var globalEnv = new(env.GlobalEnv)
	// init envName
	if len(appName) <= 0 {
		return errors.New("EnvName不能为空")
	} else {
		globalEnv.AppName = appName
		globalEnv.LoadRunEnv()
	}

	// init workdir
	if !file_utils.IsFolder(workDir) {
		return errors.New("WorkDir为空或不是一个目录")
	} else {
		globalEnv.WorkDir = workDir
	}
	if len(args) > 0 {
		globalEnv.BuildVersion = args[0]
	} else {
		globalEnv.BuildVersion = "master"
	}
	appRuntime = &runtime.AppRuntime{Env: globalEnv}
	appRuntime.Print()
	// 设置默认配置构造器
	if provider, err := cfg_provider.NewYamlConfigProvider(globalEnv); err != nil {
		return err
	} else {
		appRuntime.ConfigProvider = provider
	}
	// 设置默认日志
	if provider, err := log_provider.NewLocalFileLogProvider(globalEnv); err != nil {
		return err
	} else {
		appRuntime.LoggerProvider = provider
	}
	return nil
}
