package hospitalize

import "fmt"

type Convalesce struct {
	next Department
}

func (c *Convalesce) execute(p *Patient) {
	if p.medicineDone {
		if c.next == nil {
			fmt.Println("正常流程完成，患者康复出院")
			return
		}
		fmt.Println("患者尚未康复，观察中")
		c.next.execute(p)
		return
	}
	fmt.Println("正常流程完成，判断是否能康复出院")
	p.medicineDone = true
	c.execute(p)
}

func (c *Convalesce) setNext(next Department) Department {
	c.next = next
	return c.next
}

func (c *Convalesce) end() {
	c.next = nil
	return
}
