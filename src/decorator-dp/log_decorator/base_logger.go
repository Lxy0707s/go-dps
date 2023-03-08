package log_decorator

import "go-dps/pkg/lib/log"

/**
 * 日志装饰器,基础日志功能
 */

type BaseLogger struct {
	log log.Logger
}

var instance *BaseLogger

func NewBaseLogger(logPrefix string) LogUtil {
	if instance == nil {
		instance = &BaseLogger{log: log.NewSugar(logPrefix, false)}
	}
	return instance
}

func (b *BaseLogger) Info(message string) {
}

func (b *BaseLogger) Debug(message string) {
}

func (b *BaseLogger) Warn(message string) {
}

func (b *BaseLogger) Error(message string) {
}
