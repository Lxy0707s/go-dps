package transportation

import "fmt"

type Ship struct {
}

func (s *Ship) convey(t *Transportation, key string) {
	existNum := t.capacity
	num := 0
	if t.capacity >= 2 {
		t.capacity = t.capacity - 2
		num = 2
	} else {
		t.capacity = t.capacity - 1
		num = 1
	}
	fmt.Println("跨河路段,使用轮船运输...", "货物总数", existNum, "卸货数量", num, "剩余货物量", t.capacity)
}
