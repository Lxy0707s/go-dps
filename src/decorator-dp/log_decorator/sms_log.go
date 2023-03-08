package log_decorator

import "go-dps/pkg/lib/log"

type SMSLog struct {
	log    log.Logger
	logger LogUtil
}

var smsInstance *SMSLog

func NewSMSLog(logPrefix string, baseLogger LogUtil) LogUtil {
	if smsInstance == nil {
		smsInstance = &SMSLog{
			log:    log.NewSugar(logPrefix, false),
			logger: baseLogger,
		}
	}
	return smsInstance
}

func (s *SMSLog) Info(message string) {
	if s.logger != nil {
		s.logger.Info(message)
	}
	s.log.Info(message)
}

func (s *SMSLog) Debug(message string) {
	if s.logger != nil {
		s.logger.Debug(message)
	}
	s.log.Debug(message)
}

func (s *SMSLog) Warn(message string) {
	if s.logger != nil {
		s.logger.Warn(message)
	}
	s.log.Warn(message)
}

func (s *SMSLog) Error(message string) {
	if s.logger != nil {
		s.logger.Error(message)
	}
	s.log.Error(message)
}
