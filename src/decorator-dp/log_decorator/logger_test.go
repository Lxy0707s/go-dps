package log_decorator

import (
	"testing"
)

func Test(t *testing.T) {
	logger := NewEmailLog("Email", NewSMSLog("SMS", nil))
	logger.Info("hello")
	logger.Warn("warn")
	logger.Error("error，checking please")
}
