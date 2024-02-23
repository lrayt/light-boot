package ip

import (
	"net"
	"strings"
)

// GetLocalHostIP 获取本机IP
func GetLocalHostIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "127.0.0.1"
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return strings.Split(localAddr.String(), ":")[0]
}

func PortInUse(port uint32) bool {
	// todo: 端口是否占用
	//conn, err := net.DialTimeout("tcp", net.JoinHostPort("", port), time.Second)
	//if err != nil {
	//	return false
	//}
	//conn.Close()
	return false
}

func GetMacAddressByNet() ([]string, error) {
	var adapterList []string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		adapterList = append(adapterList, macAddr)
	}
	return adapterList, nil
}
