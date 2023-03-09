package produce

import "fmt"

// Delivered 交货
type Delivered struct {
	next  FactoryProcess
	model string
}

func (d *Delivered) qcQa(m *Material) bool {
	return true
}

func (d *Delivered) NextProcess(f FactoryProcess, strictType string) FactoryProcess {
	if f != nil {
		d.next = f
		d.next.withModel(strictType)
	}
	return d.next
}

func (d *Delivered) withModel(model string) {
	d.model = model
	return
}

func (d *Delivered) produce(m *Material) {
	m.SetTCO(10)
	if d.next == nil {
		fmt.Println(m)
		d.End()
		return
	}
	d.next.Execute(m)
}

func (d *Delivered) Execute(m *Material) {
	switch d.model {
	case Normal:
		fmt.Println("允许运输质量异常")
		m.SetDescription("允许运输质量异常")
	case Strict:
		if !d.qcQa(m) {
			m.SetQuarantineLabel(false)
			d.End()
			return
		}
	}
	d.produce(m)
}

func (d *Delivered) End() {
	d.next = nil
	return
}
