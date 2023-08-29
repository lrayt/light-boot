package http_manager

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lrayt/light-boot/convention"
	"github.com/lrayt/light-boot/core"
	"net/http"
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
