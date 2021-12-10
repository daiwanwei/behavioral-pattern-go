package shops

import (
	"fmt"
	"pattern-golang/behavioral/strategy/logistics"
)

type SportShop struct {
	logistic logistics.Logistic
}

func NewSportShop(logistic logistics.Logistic) *SportShop {
	return &SportShop{
		logistic: logistic,
	}
}

func (shop *SportShop) Buy(buyer, product string) {
	fmt.Printf("buy %s", product)
	shop.logistic.SendTo(buyer)
}
