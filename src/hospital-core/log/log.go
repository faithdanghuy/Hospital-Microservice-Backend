package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

var (
	logger  *zap.Logger
	syncOne sync.Once
)

func Must(logger *zap.Logger, err error) *zap.Logger {
	if err != nil {
		panic(err)
	}
	return logger
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(level.CapitalString())
	fmt.Println("fsdfsf")
}

func init() {
	syncOne.Do(func() {
		cfg := zap.Config{
			Encoding:    "json",
			Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
			OutputPaths: []string{"stderr"},

			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:   "message",
				TimeKey:      "time",
				LevelKey:     "level",
				CallerKey:    "caller",
				EncodeCaller: zapcore.ShortCallerEncoder,
				EncodeLevel:  CustomLevelEncoder,
				EncodeTime:   SyslogTimeEncoder,
			},
		}

		logger = Must(cfg.Build(zap.AddCaller(), zap.AddCallerSkip(1)))
	})
}

func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}

func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	logger.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}

func Panic(message string, fields ...zap.Field) {
	logger.Panic(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	logger.Fatal(message, fields...)
}

func Debugf(message string, args ...interface{}) {
	logger.Sugar().Debugf(message, args...)
}

func Infof(message string, args ...interface{}) {
	logger.Sugar().Infof(message, args...)
}

func Warnf(message string, args ...interface{}) {
	logger.Sugar().Warnf(message, args...)
}

func Errorf(message string, args ...interface{}) {
	logger.Sugar().Errorf(message, args...)
}

func Panicf(message string, args ...interface{}) {
	logger.Sugar().Panicf(message, args...)
}

func Fatalf(message string, args ...interface{}) {
	logger.Sugar().Fatalf(message, args...)
}
