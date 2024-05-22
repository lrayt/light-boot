package convention

import (
	"errors"
	"github.com/lrayt/light-boot/core"
	"github.com/lrayt/light-boot/pkg/file_utils"
	"strconv"
	"strings"
)

type StaticMap struct {
	Route    string
	FilePath string
}

type HttpConf struct {
	Host       string
	Port       uint32
	UseTLS     bool
	Pem        string
	Key        string
	Static     []StaticMap
	EnableCORS bool
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

// 解析Static
func (c HttpConf) StaticParser() error {
	for _, o := range c.Static {
		o.FilePath = strings.Replace(o.FilePath, "${WorkDir}", core.GWorkDir(), 1)
		if !file_utils.IsFolder(o.FilePath) {
			return errors.New("静态资源路径，不存在:" + o.FilePath)
		}
	}
	return nil
}

func (c HttpConf) HasStatic(route string) bool {
	for _, o := range c.Static {
		if o.Route == route {
			o.FilePath = strings.Replace(o.FilePath, "${WorkDir}", core.GWorkDir(), 1)
			return file_utils.IsFolder(o.FilePath)
		}
	}
	return false
}

func (c HttpConf) GetStaticPath(route string) (string, error) {
	for _, o := range c.Static {
		if o.Route != route {
			continue
		}
		var targetPath = strings.Replace(o.FilePath, "${WorkDir}", core.GWorkDir(), 1)
		if file_utils.IsFolder(targetPath) {
			return targetPath, nil
		}
	}
	return "", errors.New("未找到静态资源配置")
}
