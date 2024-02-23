package ip

import "testing"

func TestGetMacAddressByNet(t *testing.T) {
	macList, err := GetMacAddressByNet()
	t.Log(macList)
	t.Log(err)
}
