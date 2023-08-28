package convention

type HttpConf struct {
	Port   uint32
	UseTLS bool
	Pem    string
	Key    string
}
