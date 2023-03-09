package hospitalize

import "fmt"

type Prescribe struct {
	next Department
}

func (pr *Prescribe) execute(p *Patient) {
	if p.medicineDone {
		if pr.next == nil {
			fmt.Println("正常流程完成，患者离开")
			return
		}
		fmt.Println("医生已经给病人开完药了")
		pr.next.execute(p)
		return
	}
	fmt.Println("医生正在开药", "病人:", p.name)
	p.medicineDone = true
	pr.execute(p)
}

func (pr *Prescribe) setNext(next Department) Department {
	pr.next = next
	return pr.next
}

func (pr *Prescribe) end() {
	pr.next = nil
	return
}
