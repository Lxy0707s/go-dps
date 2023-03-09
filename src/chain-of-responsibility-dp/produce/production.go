package produce

import "fmt"

// Production 生产加工
type Production struct {
	next  FactoryProcess
	model string
}

func (p *Production) qcQa(m *Material) bool {
	return true
}

func (p *Production) withModel(model string) {
	p.model = model
	return
}

func (p *Production) produce(m *Material) {
	m.SetTCO(500)
	if p.next == nil {
		p.End()
		return
	}
	p.next.Execute(m)
}

func (p *Production) NextProcess(f FactoryProcess, strictType string) FactoryProcess {
	p.next = nil
	if f != nil {
		p.next = f
		p.next.withModel(strictType)
	}
	return p.next
}

func (p *Production) Execute(m *Material) {
	switch p.model {
	case Normal:
		fmt.Println("允许生产质量异常")
		m.SetDescription("允许生产质量异常")
	case Strict:
		if !p.qcQa(m) {
			m.SetQuarantineLabel(false)
			p.End()
			return
		}
	}
	p.produce(m)
}

func (p *Production) End() {
	p.next = nil
	return
}
