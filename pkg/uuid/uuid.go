package uuid

import (
	"fmt"
	"github.com/rs/xid"
)

func GenUUID() string {
	return xid.New().String()
}

func GenUUIDWithPrefix(prefix string) string {
	return fmt.Sprintf("%s%s", prefix, xid.New().String())
}

func IsIP(ip string) bool {
	return len(ip) > 0
}

func IsPort(port uint32) bool {
	return port < 65535 && port > 0
}
