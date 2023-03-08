package abstract_factory

import "fmt"

// IFactory 抽象工厂，包含一系列创建产品的抽象
type IFactory interface {
	makeGoodX(goodName ...string) []GoodsAction1
	makeGoodY(goodName ...string) []GoodsAction2
}

var (
	FactoryA = "factory-a"
	FactoryB = "factory-b"
)

func GetFactory(tp string) IFactory {
	switch tp {
	case FactoryA:
		return &A{}
	case FactoryB:
		return &B{}
	default:
		fmt.Println("not found type")
	}
	return nil
}
