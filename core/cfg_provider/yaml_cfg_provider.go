package cfg_provider

import (
	"errors"
	"fmt"
	"github.com/lrayt/light-boot/core/env"
	"github.com/spf13/viper"
	"path/filepath"
)

func NewYamlConfigProvider(globalEnv *env.GlobalEnv) (*YamlConfigProvider, error) {
	var cfgFile = filepath.Join(globalEnv.WorkDir, "configs", fmt.Sprintf("skeleton-%s-conf.yaml", globalEnv.RunEnv))
	var cfgContainer = &YamlConfigProvider{medium: viper.New()}
	cfgContainer.medium.SetConfigType("yaml")
	cfgContainer.medium.SetConfigFile(cfgFile)
	if err := cfgContainer.medium.ReadInConfig(); err != nil {
		return nil, errors.New(fmt.Sprintf("加载配置文件[%s]失败,err: %s", cfgFile, err))
	} else {
		fmt.Printf("加载配置文件[%s]\n", cfgFile)
	}
	return cfgContainer, nil
}

type YamlConfigProvider struct {
	medium *viper.Viper
}

func (y YamlConfigProvider) GetValue(key string) interface{} {
	return y.medium.Get(key)
}

func (y YamlConfigProvider) PackConf(cfgId string, obj interface{}) error {
	return y.medium.UnmarshalKey(cfgId, obj)
}

func (y YamlConfigProvider) PackConfToMap(cfgId string) map[string]interface{} {
	return y.medium.GetStringMap(cfgId)
}

func (y YamlConfigProvider) GetIntSlice(key string) []int {
	return y.medium.GetIntSlice(key)
}
