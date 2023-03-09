package hospitalize

import "fmt"

type Convalesce struct {
	next Process
}

func (c *Convalesce) execute(p *Patient) {
	if p.convalesceDone {
		if c.next == nil {
			fmt.Println("正常流程完成，患者康复出院")
			return
		}
		fmt.Println("患者尚未康复，观察中")
		c.next.execute(p)
		return
	}
	fmt.Println("正常流程完成，判断是否能康复出院")
	p.convalesceDone = true
	c.execute(p)
}

func (c *Convalesce) setNext(next Process) Process {
	c.next = next
	return c.next
}

func (c *Convalesce) end() {
	c.next = nil
	return
}
