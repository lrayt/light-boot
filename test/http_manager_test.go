package test

import (
	"context"
	"fmt"
	"github.com/lrayt/light-boot/core"
	"github.com/lrayt/light-boot/kit/http_manager"
	"log"
	"path/filepath"
	"reflect"
	"testing"
)

func init() {
	rootPath, pathErr := filepath.Abs("")
	if pathErr != nil {
		log.Fatalf("获取项目工作路径失败,%s\n", pathErr.Error())
	}

	if err := core.InitEnv("light-boot", rootPath, "0.1.5"); err != nil {
		log.Fatalf("init app err:%s\n", err.Error())
	}
}

type UserInfo struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func CreateUser(ctx context.Context, req *UserInfo) interface{} {
	fmt.Println("===<<<")
	//fmt.Printf("======>%s====>%d\n", req.Name, req.Age)
	lg := core.NewLoggerWithCTX("CreateUser", ctx)
	lg.Info("hello", map[string]interface{}{"req": req})
	return &UserInfo{
		Name: "abc",
		Age:  22,
	}
}

func TestGinHttpProvider(t *testing.T) {
	p := http_manager.NewGinHttpProvider()
	group := p.Group("api/v1")
	group.Post("/create", CreateUser)
	err := p.Run()
	t.Log(err)
}

type User struct {
	Name string `json:"name"`
}

func NewObj(mapData map[string]interface{}, target interface{}) interface{} {
	targetType := reflect.TypeOf(target)
	if targetType.Kind() != reflect.Ptr || targetType.Elem().Kind() != reflect.Struct {
		return nil
	}
	targetElem := targetType.Elem()
	req := reflect.New(targetElem).Elem()
	for i := 0; i < targetElem.NumField(); i++ {
		field := targetElem.Field(i)
		tag := field.Tag.Get("json")
		fmt.Println("===>", tag)
		if tag == "" {
			continue
		}

		fieldValue, found := mapData[tag]
		if !found {
			return fmt.Errorf("missing key '%s' in map data", tag)
		}

		structField := req.Field(i)
		if !structField.CanSet() {
			return fmt.Errorf("cannot set value for field '%s'", field.Name)
		}

		value := reflect.ValueOf(fieldValue)
		if value.Type().ConvertibleTo(structField.Type()) {
			structField.Set(value.Convert(structField.Type()))
		} else {
			return fmt.Errorf("cannot convert value for field '%s'", field.Name)
		}
	}
	return req
}

func TestReflect(t *testing.T) {
	data := map[string]interface{}{"name": "lirui"}
	o := NewObj(data, &User{})
	obj, ok := o.(*User)
	t.Log(o, obj, ok)
}
