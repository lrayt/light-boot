package cfg_provider

type ConfigProvider interface {
	GetValue(key string) interface{}
	PackConf(cfgId string, obj interface{}) error
	PackConfToMap(cfgId string) map[string]interface{}
	GetIntSlice(key string) []int
}
