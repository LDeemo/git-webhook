package common

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/**
按照500M文件大小切割文件
最多保留200个文件
最多保留30天
支持文件压缩
*/

var logger = &Logger{}

type Logger struct {
	l *zap.SugaredLogger
}

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

//Filename: 日志文件的位置
//MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
//MaxBackups：保留旧文件的最大个数
//MaxAges：保留旧文件的最大天数
//Compress：是否压缩/归档旧文件

// InitLogger 初始化logger
func InitLogger() {
	fileName := "./logs/console.log"
	level := getLoggerLevel("debug")
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    500,
		MaxBackups: 200,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   true,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))
	tmpLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger.l = tmpLogger.Sugar()
}

// GetLogger 获取logger
func GetLogger() *Logger {
	if logger != nil {
		return logger
	}
	return nil
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func (logger *Logger) Debug(requestID, template string, args ...interface{}) {
	logger.l.Debugw(fmt.Sprintf(template, args...), zap.String("requestID", requestID))
}

func (logger *Logger) Info(requestID, template string, args ...interface{}) {
	logger.l.Infow(fmt.Sprintf(template, args...), zap.String("requestID", requestID))
}

func (logger *Logger) Warn(requestID, template string, args ...interface{}) {
	logger.l.Warnw(fmt.Sprintf(template, args...), zap.String("requestID", requestID))
}

func (logger *Logger) Error(requestID, template string, args ...interface{}) {
	logger.l.Errorw(fmt.Sprintf(template, args...), zap.String("requestID", requestID))
}

func (logger *Logger) DPanic(requestID, template string, args ...interface{}) {
	logger.l.DPanicw(fmt.Sprintf(template, args...), zap.String("requestID", requestID))
}

func (logger *Logger) Panic(requestID, template string, args ...interface{}) {
	logger.l.Panicw(fmt.Sprintf(template, args...), zap.String("requestID", requestID))
}

func (logger *Logger) Fatal(requestID, template string, args ...interface{}) {
	logger.l.Fatalw(fmt.Sprintf(template, args...), zap.String("requestID", requestID))
}
