package main

import (
	"behavioral-pattern-go/interpreter/calculator"
	"fmt"
)

func main() {
	c := calculator.Calculator{}
	res := c.Calculate("10 2 + 5 -")
	fmt.Println(res)
}
