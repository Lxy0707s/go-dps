package produce

import "fmt"

// Packaging 生产包装
type Packaging struct {
	next  FactoryProcess
	model string
}

func (pack *Packaging) qcQa(m *Material) bool {
	return true
}

func (pack *Packaging) withModel(model string) {
	pack.model = model
	return
}

func (pack *Packaging) produce(m *Material) {
	m.SetTCO(80)
	if pack.next == nil {
		pack.End()
		return
	}
	pack.next.Execute(m)
}

func (pack *Packaging) NextProcess(f FactoryProcess, strictType string) FactoryProcess {
	if f != nil {
		pack.next = f
		pack.next.withModel(strictType)
	}
	return pack.next
}

func (pack *Packaging) Execute(m *Material) {
	switch pack.model {
	case Normal:
		fmt.Println("允许包装质量异常")
		m.SetDescription("允许包装质量异常")
	case Strict:
		if !pack.qcQa(m) {
			m.SetQuarantineLabel(false)
			pack.End()
			return
		}
	}
	pack.produce(m)
}

func (pack *Packaging) End() {
	pack.next = nil
	return
}
