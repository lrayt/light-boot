package rpc_manager

import (
	"github.com/lrayt/light-boot/convention"
	"github.com/lrayt/light-boot/core"
)

func GetRPCConf() (*convention.RPCConf, error) {
	var rpcConf = new(convention.RPCConf)
	var err = core.GConfigs().PackConf("rpc", rpcConf)
	return rpcConf, err
}
