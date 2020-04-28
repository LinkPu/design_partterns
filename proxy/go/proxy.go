package main

import "fmt"

// subject
type Specialty interface {
	Pay(money int)
}

// real subject
type SiChuanSpecialty struct{}

var realSiChuanSpecialty = SiChuanSpecialty{}

var _ Specialty = SiChuanSpecialty{}

func (SiChuanSpecialty) Pay(money int) {
	fmt.Println("Pay SiChuan specialty: ￥", money)
}

// proxy
type TaobaoProxy struct{}

var taobaoProxy = TaobaoProxy{}

var _ Specialty = TaobaoProxy{}

func (TaobaoProxy) Pay(money int) {
	fmt.Println("Pay Taobao proxy: ￥", money)
	money = taobaoProxy.preRequest(money)
	realSiChuanSpecialty.Pay(money)
	taobaoProxy.postRequest()
}

func (TaobaoProxy) preRequest(money int) int {
	// charge agency fee
	fmt.Println("Charge agency fee: ￥5.00")
	return money - 5
}

func (TaobaoProxy) postRequest() {
	fmt.Println("Product delivery")
}

func main() {
	proxy := TaobaoProxy{}
	proxy.Pay(100)
}
