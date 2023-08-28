package test

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lrayt/light-boot/core"
	"github.com/lrayt/light-boot/kit/http_manager"
	"log"
	"path/filepath"
	"testing"
)

func init() {
	rootPath, pathErr := filepath.Abs("")
	if pathErr != nil {
		log.Fatalf("获取项目工作路径失败,%s\n", pathErr.Error())
	}

	if err := core.InitApp("light-boot", rootPath, "0.1.5"); err != nil {
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
	lg.Info("hello")
	return gin.H{"code": 11100}
}

func TestGinHttpProvider(t *testing.T) {
	p := http_manager.NewGinHttpProvider()
	group := p.Group("api/v1")
	group.Post("/create", CreateUser)
	err := p.Run()
	t.Log(err)
}
