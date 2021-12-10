package main

import (
	"behavioral-pattern-go/strategy/logistics"
	"behavioral-pattern-go/strategy/shops"
)

func main() {
	car := logistics.Car{}
	shop := shops.NewSportShop(&car)
	shop.Buy("ann", "t-shirt")
}
