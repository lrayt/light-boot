package log_provider

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/lrayt/light-boot/core/env"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"time"
)

type LocalFileLogProvider struct {
	logger *logrus.Logger
}

func NewLocalFileLogProvider(env *env.GlobalEnv) (*LocalFileLogProvider, error) {
	var logger = logrus.New()
	// 日志格式
	logger.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		FullTimestamp:   true,
		ForceQuote:      true,
	}
	var file = filepath.Join(env.WorkDir, "logs", env.AppName+"-%Y%m%d.log")
	var writer, err = rotatelogs.New(file, rotatelogs.WithRotationTime(24*time.Hour), rotatelogs.WithRotationCount(7))
	if err != nil {
		return nil, err
	}
	lfHook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.PanicLevel: writer,
		},
		&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05.000",
			ForceColors:     true,
			FullTimestamp:   true,
			ForceQuote:      true,
		})
	logger.AddHook(lfHook)
	return &LocalFileLogProvider{logger: logger}, nil
}

func (p LocalFileLogProvider) NewLogger(commonField map[string]interface{}) Logger {
	return &LocalFileLogLogger{Logger: p.logger, CommonFields: commonField}
}

type LocalFileLogLogger struct {
	Logger       *logrus.Logger
	CommonFields map[string]interface{}
}

func (s LocalFileLogLogger) Info(msg string, args ...map[string]interface{}) {
	var fields = s.CommonFields
	if len(args) > 0 && len(args[0]) > 0 {
		for k, v := range args[0] {
			fields[k] = v
		}
	}
	s.Logger.WithFields(fields).Info(msg)
}

func (s LocalFileLogLogger) Success(msg string, args ...map[string]interface{}) string {
	var fields = s.CommonFields
	if len(args) > 0 && len(args[0]) > 0 {
		for k, v := range args[0] {
			fields[k] = v
		}
	}
	s.Logger.WithFields(fields).Info(msg)
	return "Success"
}

func (s LocalFileLogLogger) Warn(msg string, args ...map[string]interface{}) {
	var fields = s.CommonFields
	if len(args) > 0 && len(args[0]) > 0 {
		for k, v := range args[0] {
			fields[k] = v
		}
	}
	s.Logger.WithFields(fields).Warn(msg)
}

func (s LocalFileLogLogger) Error(msg string, args ...map[string]interface{}) {
	var fields = s.CommonFields
	if len(args) > 0 && len(args[0]) > 0 {
		for k, v := range args[0] {
			fields[k] = v
		}
	}
	s.Logger.WithFields(fields).Error(msg)
}

func (s LocalFileLogLogger) ErrorF(msg string, err error, args ...map[string]interface{}) {
	var fields = s.CommonFields
	if len(args) > 0 && len(args[0]) > 0 {
		for k, v := range args[0] {
			fields[k] = v
		}
	}
	if err != nil {
		msg += ",err:" + err.Error()
	}
	s.Logger.WithFields(fields).Error(msg)
}

func (s LocalFileLogLogger) NewError(msg string, args ...map[string]interface{}) error {
	var fields = s.CommonFields
	if len(args) > 0 && len(args[0]) > 0 {
		for k, v := range args[0] {
			fields[k] = v
		}
	}
	s.Logger.WithFields(fields).Error(msg)
	return errors.New(msg)
}

func (s LocalFileLogLogger) NewErrorF(msg string, err error, args ...map[string]interface{}) error {
	var fields = s.CommonFields
	if len(args) > 0 && len(args[0]) > 0 {
		for k, v := range args[0] {
			fields[k] = v
		}
	}
	if err != nil {
		msg += ",err:" + err.Error()
	}
	s.Logger.WithFields(fields).Error(msg)
	return errors.New(msg)
}
