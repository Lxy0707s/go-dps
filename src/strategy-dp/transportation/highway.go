package transportation

import "fmt"

type Highway struct {
}

func (h *Highway) convey(t *Transportation, key string) {
	existNum := t.capacity
	num := 0
	if t.capacity >= 4 {
		t.capacity = t.capacity - 4
		num = 4
	} else {
		t.capacity = t.capacity - 1
		num = 1
	}
	fmt.Println("本省路段,使用公路运输...", "货物总数", existNum, "卸货数量", num, "剩余货物量", t.capacity)
}
