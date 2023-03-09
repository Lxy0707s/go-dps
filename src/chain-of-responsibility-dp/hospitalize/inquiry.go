package hospitalize

import "fmt"

type Inquiry struct {
	next Process
}

func (in *Inquiry) execute(p *Patient) {
	if p.checkUpDone {
		fmt.Println("医生问诊完成")
		if in.next == nil {
			fmt.Println("来访者健康，无需医治")
			return
		}
		in.next.execute(p)
		return
	}
	fmt.Println("医生正在问诊")
	p.checkUpDone = true
	in.execute(p)
}

func (in *Inquiry) setNext(next Process) Process {
	in.next = next
	return in.next
}

func (in *Inquiry) end() {
	in.next = nil
	return
}
