package http_manager

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lrayt/light-boot/convention"
	"github.com/lrayt/light-boot/core"
	"reflect"
	"strings"
)

type GinRouterGroup struct {
	GroupName string
	Engine    *gin.Engine
}

func (g GinRouterGroup) Get(route string, handler RouterHandler) {

}

func (g GinRouterGroup) Post(route string, handler RouterHandler) {
	route = g.GroupName + "/" + strings.Trim(route, "/")
	g.Engine.POST(route, func(c *gin.Context) {
		handlerType := reflect.TypeOf(handler)
		// check handler type
		if handlerType.Kind() != reflect.Func {
			return
		}
		// check handler param num
		if handlerType.NumIn() != 2 {
			return
		}
		// check handler first param
		if handlerType.In(0) != reflect.TypeOf((*context.Context)(nil)).Elem() {
			return
		}
		// check handler last param
		if handlerType.In(1).Kind() != reflect.Ptr {
			return
		}
		// fill req
		// req
		req := reflect.New(handlerType.In(1)).Elem()
		if err := c.Bind(req); err != nil {
			return
		}
		handlerValue := reflect.ValueOf(handler)
		args := []reflect.Value{reflect.ValueOf(ToCTX(c)), req}
		handlerValue.Call(args)
	})
}

type GinHttpProvider struct {
	Engine *gin.Engine
}

func NewGinHttpProvider() *GinHttpProvider {
	r := gin.Default()
	r.Use(CORSMiddleware())
	return &GinHttpProvider{
		Engine: r,
	}
}

func (p GinHttpProvider) Group(name string) RouterGroup {
	return &GinRouterGroup{
		GroupName: strings.Trim(name, "/"),
		Engine:    p.Engine,
	}
}

func (p GinHttpProvider) Run() error {
	var cfg = new(convention.HttpConf)
	if err := core.GConfigs().PackConf("http", cfg); err != nil {
		return err
	}
	return p.Engine.Run(fmt.Sprintf("0.0.0.0:%d", cfg.Port))
}
