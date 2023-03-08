package light

import "go-dps/pkg/lib/log"

type Incandescent struct {
	log   log.Logger
	light Light
}

func (in *Incandescent) Charge() Light {
	in.log.Info("需要接通电源")
	return in
}

func (in *Incandescent) Shine() {
	in.log.Info("发白色光")
}
