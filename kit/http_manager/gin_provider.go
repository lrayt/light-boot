package http_manager

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lrayt/light-boot/convention"
	"github.com/lrayt/light-boot/core"
	"github.com/lrayt/light-boot/core/env"
	"github.com/lrayt/light-boot/core/event_bus"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
	"time"
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
		targetType := handlerType.In(1)
		if targetType.Kind() != reflect.Ptr || targetType.Elem().Kind() != reflect.Struct {
			return
		}
		//targetElem := targetType.Elem()
		//req := reflect.New(targetElem).Elem()
		//for i := 0; i < targetElem.NumField(); i++ {
		//	field := targetElem.Field(i)
		//	tag := field.Tag.Get("json")
		//	if tag == "" {
		//		continue
		//	}
		//	c.ShouldBindBodyWith()
		//
		//	fieldValue, found := mapData[tag]
		//	if !found {
		//		return fmt.Errorf("missing key '%s' in map data", tag)
		//	}
		//
		//	structField := req.Field(i)
		//	if !structField.CanSet() {
		//		return fmt.Errorf("cannot set value for field '%s'", field.Name)
		//	}
		//
		//	value := reflect.ValueOf(fieldValue)
		//	if value.Type().ConvertibleTo(structField.Type()) {
		//		structField.Set(value.Convert(structField.Type()))
		//	} else {
		//		return fmt.Errorf("cannot convert value for field '%s'", field.Name)
		//	}
		//}
		//
		//req := reflect.New(handlerType.In(1).Elem())
		//for i := 0; i < req.NumField(); i++ {
		//	req.Field(i).SetString("lirui")
		//}

		//targetElem := handlerType.In(1).Elem()
		////targetType := targetElem.NumField()
		//
		//for i := 0; i < targetElem.NumField(); i++ {
		//	field := targetElem.Field(i)
		//	tag := field.Tag.Get("json")
		//
		//	fmt.Println("======<", tag)
		//	//if tag == "" {
		//	//	continue
		//	//}
		//	//
		//	//fieldValue, found := mapData[tag]
		//	//if !found {
		//	//	return fmt.Errorf("missing key '%s' in map data", tag)
		//	//}
		//	//
		//	//structField := targetElem.Field(i)
		//	//if !structField.CanSet() {
		//	//	return fmt.Errorf("cannot set value for field '%s'", field.Name)
		//	//}
		//	//
		//	//value := reflect.ValueOf(fieldValue)
		//	//if value.Type().ConvertibleTo(structField.Type()) {
		//	//	structField.Set(value.Convert(structField.Type()))
		//	//} else {
		//	//	return fmt.Errorf("cannot convert value for field '%s'", field.Name)
		//	//}
		//}
		//
		//fmt.Println("===>", handlerType.In(1).Elem().Kind())
		//
		//// fill req
		//// req
		req := reflect.New(targetType.Elem())
		if err := c.Bind(&req); err != nil {
			fmt.Println("---->", err)
			return
		}
		fmt.Println("req:", req)
		handlerValue := reflect.ValueOf(handler)
		args := []reflect.Value{reflect.ValueOf(ToCTX(c)), req}
		result := handlerValue.Call(args)
		fmt.Println("---<", len(result))
		c.JSON(http.StatusOK, &result[0])
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

func (p GinHttpProvider) RunWithHandler(baseUrl string, handler func(rg *gin.RouterGroup)) {
	if core.GRunEnv() == env.RunProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}
	// 读取配置
	conf, confErr := GetHttpConf()
	if confErr != nil {
		log.Fatalf("%s服务读取配置，err:%s\n", core.GAppName(), confErr.Error())
	}
	// 静态文件配置
	for _, webResource := range conf.Static {
		p.Engine.Static(webResource.Route, filepath.Join(core.GWorkDir(), webResource.FilePath))
	}
	// 跨域
	if conf.EnableCORS {
		p.Engine.Use(CORSMiddleware())
	}
	rg := p.Engine.Group(baseUrl)
	handler(rg)
	core.GEventBus().Emit(event_bus.EventBeforeHttpStart, nil)
	// 检测服务是否启动
	go func() {
		var addr = fmt.Sprintf("127.0.0.1:%d", conf.Port)
		for {
			if _, err := net.DialTimeout("tcp", addr, time.Second); err == nil {
				log.Printf("%s服务已启动，%s\n", core.GAppName(), conf.BaseUrl())
				core.GEventBus().Emit(event_bus.EventHttpStarted, nil)
				break
			}
		}
	}()
	// 服务
	srv := http.Server{
		Handler: p.Engine,
		Addr:    fmt.Sprintf("0.0.0.0:%v", conf.Port),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("%s服务启动失败，err:%s\n", core.GAppName(), err.Error())
	}
}
