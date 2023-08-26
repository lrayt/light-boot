package core

import "testing"

func TestCore(t *testing.T) {
	err := InitApp("light-boot", "E:\\workspace\\golang\\src\\github.com\\light-boot", "0.1.1")
	t.Log(err)
}
