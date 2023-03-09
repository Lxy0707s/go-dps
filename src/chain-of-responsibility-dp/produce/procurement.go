package produce

import "fmt"

// Procurement 采购
type Procurement struct {
	next  FactoryProcess
	model string
}

func (p *Procurement) qcQa(m *Material) bool {
	return true
}

func (p *Procurement) withModel(model string) {
	p.model = model
	return
}

func (p *Procurement) produce(m *Material) {
	m.SetTCO(1000)
	if p.next == nil {
		p.End()
		return
	}
	p.next.Execute(m)
}

func (p *Procurement) NextProcess(f FactoryProcess, strictType string) FactoryProcess {
	if f != nil {
		p.next = f
		p.next.withModel(strictType)
	}
	return p.next
}

func (p *Procurement) Execute(m *Material) {
	switch p.model {
	case Normal:
		fmt.Println("允许采购质量异常")
		m.SetDescription("允许采购质量异常")
	case Strict:
		if !p.qcQa(m) {
			m.SetQuarantineLabel(false)
			p.End()
			return
		}
	}
	p.produce(m)
}

func (p *Procurement) End() {
	p.next = nil
	return
}
