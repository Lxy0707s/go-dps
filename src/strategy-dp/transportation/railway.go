package transportation

import "fmt"

type Railway struct {
}

func (r *Railway) convey(t *Transportation, key string) {
	existNum := t.capacity
	num := 0
	if t.capacity >= 3 {
		t.capacity = t.capacity - 3
		num = 3
	} else {
		t.capacity = t.capacity - 1
		num = 1
	}
	fmt.Println("远程路段,使用铁路运输...", "货物总数", existNum, "卸货数量", num, "剩余货物量", t.capacity)
}
