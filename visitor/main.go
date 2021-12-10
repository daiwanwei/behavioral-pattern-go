package main

import "behavioral-pattern-go/visitor/museums"

func main() {
	museum := museums.NewMuseum()
	man := American{
		"USA",
	}
	museum.West.Welcome(&man)
	museum.East.Welcome(&man)

	man2 := Asian{
		"TW",
	}
	museum.West.Welcome(&man2)
	museum.East.Welcome(&man2)
}
