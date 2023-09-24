package convention

type StaticMap struct {
	Route    string
	FilePath string
}

type HttpConf struct {
	Port   uint32
	UseTLS bool
	Pem    string
	Key    string
	Static []StaticMap
}
