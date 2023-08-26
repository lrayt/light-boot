package ts_error

import "fmt"

// 统一错误码
const (
	Success      = 0
	ParamInvalid = 10001 // 无效的参数
	ParamIsEmpty = 10002

	DataFromDBErr = 20001 // FROM数据库错误
	DataToDBErr   = 20002 // TO数据库错误

	PortHoldErr      = 30001
	DoHttpRequestErr = 30002
	SystemErr        = 40001

	BusinessErr = 50001
)

type BaseResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

var BindErr = &BaseResponse{Code: ParamInvalid, Msg: "无效的参数"}

var EmptyParamErr = func(paramName string) *BaseResponse {
	return &BaseResponse{Code: ParamIsEmpty, Msg: fmt.Sprintf("参数[%s]不能为空", paramName)}
}
