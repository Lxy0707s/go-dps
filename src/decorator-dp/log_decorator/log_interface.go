package log_decorator

type (
	LogUtil interface {
		Debug(message string)
		Info(message string)
		Warn(message string)
		Error(message string)
	}
)
