package hospitalize

import "fmt"

// Reception 接待
type Reception struct {
	next Process
}

func (r *Reception) execute(p *Patient) {
	if p.receptionDone {
		if r.next == nil {
			fmt.Println("来访者接待完成,不是患者，不进入医患流程")
			return
		}
		fmt.Println("来访者接待完成，需要就医，进入下一流程")
		r.next.execute(p)
		return
	}
	fmt.Println("正在接待来访者")
	p.receptionDone = true
	r.execute(p)
}

func (r *Reception) setNext(next Process) Process {
	r.next = next
	return r.next
}

func (r *Reception) end() {
	r.next = nil
	return
}
