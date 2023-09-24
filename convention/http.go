package convention

import "strconv"

type StaticMap struct {
	Route    string
	FilePath string
}

type HttpConf struct {
	Host   string
	Port   uint32
	UseTLS bool
	Pem    string
	Key    string
	Static []StaticMap
}

func (c HttpConf) BaseUrl() string {
	if len(c.Host) <= 0 {
		c.Host = "127.0.0.1"
	}
	var url string
	if c.UseTLS {
		url = "https://" + c.Host
	} else {
		url = "http://" + c.Host
	}
	if c.Port != 80 && c.Port != 443 {
		url += ":" + strconv.Itoa(int(c.Port))
	}
	return url
}
