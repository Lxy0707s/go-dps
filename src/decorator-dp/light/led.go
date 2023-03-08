package light

import "go-dps/pkg/lib/log"

type LED struct {
	log   log.Logger
	light Light
}

func (l *LED) Charge() Light {
	l.log.Info("需要接通电源")
	return l
}

func (l *LED) Shine() {
	l.log.Info("发彩色光")
}
