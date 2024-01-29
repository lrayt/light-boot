package rpc_manager

import (
	"github.com/lrayt/light-boot/core"
	"github.com/lrayt/light-boot/core/event_bus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type GRPCProvider struct {
	svr *grpc.Server
}

func NewGRPCProvider() *GRPCProvider {
	return &GRPCProvider{
		svr: grpc.NewServer(),
	}
}

func (p GRPCProvider) Setup(handler func(svr *grpc.Server)) {
	conf, getErr := GetRPCConf()
	if getErr != nil {
		log.Fatalf("rpc server run err:%s\n", getErr.Error())
	}

	listener, listenerErr := net.Listen("tcp", conf.Addr())
	if listenerErr != nil {
		log.Fatalf("rpc server listen err:%s\n", listenerErr.Error())
	}
	handler(p.svr)
	reflection.Register(p.svr)
	core.GEventBus().Emit(event_bus.EventBeforeRPCStart, nil)
	// 检测RPC是否启动
	go func() {
		for {
			if _, err2 := net.DialTimeout("tcp", conf.Addr(), time.Second); err2 == nil {
				log.Printf("%s服务已启动[%s]\n", core.GAppName(), conf.BaseUrl())
				core.GEventBus().Emit(event_bus.EventRPCStarted, nil)
				break
			}
		}
	}()

	if err := p.svr.Serve(listener); err != nil {
		log.Fatalf("%s服务启动失败，err:%s\n", core.GAppName(), err.Error())
	}
}
