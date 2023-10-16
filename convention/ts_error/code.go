package ts_error

// 统一错误码
const (
	Success      = 10000
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
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
