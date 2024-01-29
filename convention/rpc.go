package convention

import "fmt"

type RPCConf struct {
	Host string
	Port uint32
}

func (c RPCConf) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c RPCConf) BaseUrl() string {
	return fmt.Sprintf("rpc://%s:%d", c.Host, c.Port)
}
