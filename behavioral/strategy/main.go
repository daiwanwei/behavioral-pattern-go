package main

import (
	"pattern-golang/behavioral/strategy/logistics"
	"pattern-golang/behavioral/strategy/shops"
)

func main() {
	car := logistics.Car{}
	shop := shops.NewSportShop(&car)
	shop.Buy("ann", "t-shirt")
}
