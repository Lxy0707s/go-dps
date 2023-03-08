package transportation

import "fmt"

type (
	Transportation struct {
		goods       map[string][]string // 货物种类、数量
		transport   Transport           // 运输方式
		company     string              // 承运公司
		capacity    int                 // 货物运输量
		maxCapacity int                 // 可运输最大货物量
	}
	Transport interface {
		convey(t *Transportation, key string) // 具体运输方式
	}
)

func initTransportation(c string, t Transport) *Transportation {
	goods := make(map[string][]string)
	return &Transportation{
		goods:       goods,
		transport:   t,
		company:     c,
		capacity:    0,
		maxCapacity: 6,
	}
}

func (t *Transportation) setEvictionAlgo(tp Transport) {
	t.transport = tp
}

func (t *Transportation) add(key string, value []string) {
	if t.capacity >= t.maxCapacity {
		fmt.Println("货物已超载,请检查")
		t.convey(key)
		return
	}
	t.capacity = t.capacity + len(value)
	t.goods[key] = value
	t.convey(key)
}

func (t *Transportation) get() int {
	return t.capacity
}

func (t *Transportation) convey(key string) {
	t.transport.convey(t, key)
}
