package produce

import "testing"

func Test(t *testing.T) {
	procure := &Procurement{}
	product := &Production{}
	pack := &Packaging{}
	deliver := &Delivered{}

	procure.withModel(Normal)
	procure.NextProcess(product, Strict).
		NextProcess(pack, Strict).
		NextProcess(deliver, Normal).
		End()
	procure.Execute(&Material{
		Manufacturer:    "test",
		Models:          "test",
		Price:           0,
		TCO:             0,
		Name:            "产品A",
		Description:     "A-TEST",
		QuarantineLabel: true,
	})
}
