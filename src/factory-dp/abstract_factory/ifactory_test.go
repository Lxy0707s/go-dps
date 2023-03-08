package abstract_factory

import "testing"

func Test(t *testing.T) {
	// 代工厂A生产两种货物
	aFactory := GetFactory(FactoryA)
	good1A := aFactory.makeGoodX(Good1, "self-good1") //产品线1
	good2A := aFactory.makeGoodY(Good2)               //产品线2

	// 代工厂B生产两种货物
	bFactory := GetFactory(FactoryB)
	good1B := bFactory.makeGoodX(Good1)               //产品线1
	good2B := bFactory.makeGoodY(Good2, "self-good2") //产品线2

	// 模拟客户获取货物
	GetDetails(good1A, good2A, good1B, good2B) // 获取各个产品线产品
}
