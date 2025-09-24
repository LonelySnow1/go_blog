package core

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"server/global"
)

func InitLogger() *zap.Logger {
	zapConfig := global.Config.Zap
	WriteSyncer := getLogWriter(zapConfig.Filename, zapConfig.MaxSize, zapConfig.MaxBackups, zapConfig.MaxAge) //写入指定文件
	if zapConfig.IsConsolePrint {
		WriteSyncer = zapcore.NewMultiWriteSyncer(WriteSyncer, zapcore.AddSync(os.Stdout)) //多写入同步器，可以同时将日志写入到所有的参数中
	}
	encoder := getEncoder() //自定义编码器
	var coreLevel zapcore.Level
	err := coreLevel.UnmarshalText([]byte(zapConfig.Level)) // 将日志级别（如 "info"、"error"）解析为zap能识别的格式
	if err != nil {
		log.Fatalf("Failed to parse log level err: %v", err)
	}
	core := zapcore.NewCore(encoder, WriteSyncer, coreLevel) // 创建zap的核心组件(core) 三个参数分别为:编码器，写入器，日志级别
	logger := zap.New(core, zap.AddCaller())                 // 直接创建一个可用的日志实例，zap.AddCaller() ——> 自动添加调用者信息
	return logger
}
func getLogWriter(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer { // 定义一个写入器
	lumberJackLogger := &lumberjack.Logger{ // 对日志进行切割归档
		Filename:   filename,   // 这个里面就带有路径，指定文件
		MaxSize:    maxSize,    // 日志最大大小，超过会分割
		MaxBackups: maxBackups, // 保留最大个数
		MaxAge:     maxAge,     // 保留的最大天数
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder { // 为zap日志创建一个自定义的JSON格式的编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
