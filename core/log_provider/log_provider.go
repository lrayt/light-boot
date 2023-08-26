package log_provider

// LoggerProvider 日志收集器
type LoggerProvider interface {
	NewLogger(commonField map[string]interface{}) Logger
}

type Logger interface {
	Info(msg string, args ...map[string]interface{})
	Success(msg string, args ...map[string]interface{}) string
	Warn(msg string, args ...map[string]interface{})
	Error(msg string, args ...map[string]interface{})
	ErrorF(msg string, err error, args ...map[string]interface{})
	NewError(msg string, args ...map[string]interface{}) error
	NewErrorF(msg string, err error, args ...map[string]interface{}) error
}
