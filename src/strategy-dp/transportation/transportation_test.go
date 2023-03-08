package transportation

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	hw := &Highway{}
	transport := initTransportation("国际运输公司", hw)
	transport.add("a", []string{"公路运输获取1", "公路运输获取2", "公路运输获取3", "公路运输获取4"})

	air := &Air{}
	transport.setEvictionAlgo(air)
	transport.add("b", []string{"空运货物1", "空运货物2", "空运货物4", "空运货物4", "空运货物5", "空运货物6", "空运货物9"})

	ship := &Ship{}
	transport.setEvictionAlgo(ship)
	transport.add("c", []string{"水运货物1", "水运货物2"})

	rl := &Railway{}
	transport.setEvictionAlgo(rl)
	transport.add("d", []string{"铁路运输货物1", "铁路运输货物2"})

	fmt.Println("剩余货物情况", "货物量", transport.get())
}
