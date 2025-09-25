package initialize

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"os"
	"server/global"
	"server/task"
)

// ZapLogger  : 适配器 让 zap 适配 cron
// 想让cron日志记录为预期的zap格式 -> WithLogger只接受自己的log构造器 -> 必须让构造器适用原有log构造器的方法
// -> 实现原有log构造器的接口 -> 要新创建一个结构体来实现接口
type ZapLogger struct {
	logger *zap.Logger
}

// Info 在Zap中是这样定义的: func (log *zap.Logger) Info(msg string, fields ...zap.Field)
// 所以需要对keysAndValues进行结构化，让其变成zap.Field形式
func (l *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info(msg, zap.Any("keysAndValues", keysAndValues)) // zap.Any()用于结构化 keysAndValues
}
func (l *ZapLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.logger.Error(msg, zap.Any("keysAndValues", keysAndValues))
}
func NewZapLogger() *ZapLogger {
	return &ZapLogger{logger: global.Log}
}
func InitCron() {
	// 将 cron 包的日志记录转发到 zap 日志库中，实现统一的日志管理和记录
	c := cron.New(cron.WithLogger(NewZapLogger()))

	err := task.RegisterScheduledTasks(c) // 注册定时任务
	if err != nil {
		global.Log.Error("Failed to register scheduled task", zap.Error(err))
		os.Exit(1)
	}
}
