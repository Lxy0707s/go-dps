package log_decorator

import (
	"go-dps/pkg/lib/log"
)

type EmailLog struct {
	log    log.Logger
	logger LogUtil
}

var emailInstance *EmailLog

func NewEmailLog(logPrefix string, baseLogger LogUtil) LogUtil {
	if emailInstance == nil {
		emailInstance = &EmailLog{
			log:    log.NewSugar(logPrefix, false),
			logger: baseLogger,
		}
	}
	return emailInstance
}

func (e *EmailLog) Info(message string) {
	if e.logger != nil {
		e.logger.Info(message)
	}
	e.log.Info(message)
}

func (e *EmailLog) Debug(message string) {
	if e.logger != nil {
		e.logger.Debug(message)
	}
	e.log.Debug(message)
}

func (e *EmailLog) Warn(message string) {
	if e.logger != nil {
		e.logger.Warn(message)
	}
	e.log.Warn(message)
}

func (e *EmailLog) Error(message string) {
	if e.logger != nil {
		e.logger.Error(message)
	}
	e.log.Error(message)
}
