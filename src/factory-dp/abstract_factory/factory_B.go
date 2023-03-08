package abstract_factory

type B struct {
}

func (b *B) makeGoodX(goodNames ...string) []GoodsAction1 {
	var gds = make([]GoodsAction1, 0)
	for _, good := range goodNames {
		gds = append(gds, &Goods1{GoodInfo{
			name:  good + "_A",
			type_: GoodTypeX,
			price: 110,
		}})
	}
	return gds
}

func (b *B) makeGoodY(goodNames ...string) []GoodsAction2 {
	var gds = make([]GoodsAction2, 0)
	for _, good := range goodNames {
		gds = append(gds, &Goods2{GoodInfo{
			name:  good + "_A",
			type_: GoodTypeY,
			price: 110,
		}})
	}
	return gds
}
