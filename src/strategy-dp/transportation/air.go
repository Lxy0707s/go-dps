package transportation

import "fmt"

type Air struct {
}

func (a *Air) convey(t *Transportation, key string) {
	existNum := t.capacity
	num := 0
	if t.capacity >= 10 {
		t.capacity = t.capacity - 10
		num = 10
	} else {
		t.capacity = t.capacity - 1
		num = 1
	}
	fmt.Println("跨海路段,使用飞机运输...", "货物总数", existNum, "卸货数量", num, "剩余货物量", t.capacity)
}
