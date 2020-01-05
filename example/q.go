package main

import (
	"fmt"
	"github.com/Kinggo/resk/core/envelopes"
)

func main() {
	goods := envelopes.RedEnvelopeGoods{
		EnvelopeNo: "111",
	}
	r := goods
	r.EnvelopeNo = "222"
	fmt.Println(goods.EnvelopeNo)
	fmt.Println(r.EnvelopeNo)
	func() {
		fmt.Println(goods.EnvelopeNo)
		fmt.Println(r.EnvelopeNo)
	}()
}
