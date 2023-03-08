package light

import (
	"go-dps/pkg/lib/log"
	"testing"
)

func Test(t *testing.T) {
	var light Light = &Incandescent{log: log.NewSugar("Incandescent", false)}
	light.Charge().Shine()

	light = &LED{
		light: light,
		log:   log.NewSugar("LED", false),
	}
	light.Charge().Shine()
}
