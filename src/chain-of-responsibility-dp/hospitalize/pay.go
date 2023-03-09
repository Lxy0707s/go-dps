package hospitalize

import "fmt"

type Pay struct {
	next Process
}

func (pa *Pay) execute(p *Patient) {
	if p.paymentDone {
		if pa.next == nil {
			fmt.Println("无需开药，结费出院")
			return
		}
		fmt.Println("付钱完成，进入下一流程")
		pa.next.execute(p)
		return
	}
	fmt.Println("收银员正在处理交易")
	p.paymentDone = true
	pa.execute(p)
}

func (pa *Pay) setNext(next Process) Process {
	pa.next = next
	return pa.next
}

func (pa *Pay) end() {
	pa.next = nil
	return
}
