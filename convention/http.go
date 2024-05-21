package convention

import (
	"errors"
	"github.com/lrayt/light-boot/core"
	"github.com/lrayt/light-boot/pkg/file_utils"
	"path/filepath"
	"strconv"
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

func (c HttpConf) HasStatic(route string, rootPath ...string) bool {
	if len(c.Static) <= 0 {
		return false
	}
	var prefix = core.GWorkDir()
	if len(rootPath) > 0 {
		prefix = rootPath[0]
	}
	for _, o := range c.Static {
		if o.Route == route {
			return file_utils.IsFolder(filepath.Join(prefix, o.FilePath))
		}
	}
	return false
}

func (c HttpConf) GetStaticPath(route string, rootPath ...string) (string, error) {
	var prefix = core.GWorkDir()
	if len(rootPath) > 0 {
		prefix = rootPath[0]
	}
	for _, o := range c.Static {
		if o.Route != route {
			continue
		}
		var targetPath = filepath.Join(prefix, o.FilePath)
		if file_utils.IsFolder(targetPath) {
			return targetPath, nil
		}
	}
	return "", errors.New("未找到静态资源配置")
}
