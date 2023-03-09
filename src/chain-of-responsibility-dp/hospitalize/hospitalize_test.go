package hospitalize

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	reception := &Reception{}
	inquiry := &Inquiry{}
	prescribe := &Prescribe{}
	cashier := &Pay{}
	convalesce := &Convalesce{}

	fmt.Println("-----------------------------")
	// 接待 	问诊 付钱 开药 出院判断 进入观察复诊 出院
	reception.setNext(inquiry).setNext(cashier).setNext(prescribe).setNext(convalesce).setNext(&Inquiry{}).end()
	reception.execute(&Patient{name: "病人小雷"})
	fmt.Println("-----------------------------")
	// 紧急情况:跳过接待，	问诊 付钱 开药 出院判断
	inquiry.setNext(cashier).setNext(prescribe).setNext(convalesce).end()
	inquiry.execute(&Patient{name: "急诊病人小张"})
	fmt.Println("-----------------------------")
}
