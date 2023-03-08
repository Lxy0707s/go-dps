package abstract_factory

import "fmt"

// 抽象产品
type (
	GoodType interface {
		setName(name string)
		setPrice(price float64)
		getName() string
		getPrice() float64
	}
	GoodInfo struct {
		name  string
		type_ string
		price float64
	}
	// Goods1 Goods2 产品种类
	Goods1 struct {
		GoodInfo
	}
	Goods2 struct {
		GoodInfo
	}

	// GoodsAction1 GoodsAction2 产品操作
	GoodsAction1 interface {
		GoodType
	}
	GoodsAction2 interface {
		GoodType
	}
)

var (
	Good1     = "good1"
	Good2     = "good2"
	GoodTypeX = "goodx"
	GoodTypeY = "goody"
)

func (g1 *Goods1) setName(name string) {
	g1.name = name
}

func (g1 *Goods1) getName() string {
	return g1.name
}

func (g1 *Goods1) setPrice(price float64) {
	g1.price = price
}

func (g1 *Goods1) getPrice() float64 {
	return g1.price
}

/*----------------分割线-------------------*/

func (g2 *Goods2) setName(name string) {
	g2.name = name
}

func (g2 *Goods2) getName() string {
	return g2.name
}

func (g2 *Goods2) setPrice(price float64) {
	g2.price = price
}

func (g2 *Goods2) getPrice() float64 {
	return g2.price
}

/*----------------分割线-------------------*/

func GetDetails(goodTypes ...interface{}) {
	for _, type_ := range goodTypes {
		switch type_.(type) {
		case []GoodsAction1:
			goodAcs, ok := type_.([]GoodsAction1)
			if !ok {
				return
			}
			for _, ac := range goodAcs {
				fmt.Printf("goods: %s", ac.getName())
				fmt.Printf(", price: %.2f", ac.getPrice())
				fmt.Println()
				fmt.Println("---------------------------")
			}
		case []GoodsAction2:
			goodAcs, ok := type_.([]GoodsAction2)
			if !ok {
				return
			}
			for _, ac := range goodAcs {
				fmt.Printf("goods: %s", ac.getName())
				fmt.Printf(", price: %.2f", ac.getPrice())
				fmt.Println()
				fmt.Println("---------------------------")
			}
		}
	}
}
