package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger

func init() {
	// 自定义时间格式
	customTimeEncoder := zapcore.TimeEncoderOfLayout("2006/01/02 - 15:04:05")
	// 定义EncoderConfig
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写级别编码器
		EncodeTime:     customTimeEncoder,             // 自定义时间编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短调用者编码器
	}
	// 创建一个Encoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	// 创建一个Core，并设置日志级别为Info
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.NewAtomicLevelAt(zap.InfoLevel))
	// 创建Logger
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}

func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func Error(err error, fields ...zap.Field) {
	logger.Error(err.Error(), fields...)
}
