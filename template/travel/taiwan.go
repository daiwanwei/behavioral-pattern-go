package travel

import (
	"fmt"
)

type TaiwanTravel struct {
	Travel
}

func NewTaiwanTravel() *TaiwanTravel {
	t := &TaiwanTravel{}
	tmpl := Travel{custom: t}
	t.Travel = tmpl
	return t
}

func (travel *TaiwanTravel) StartOff() {
	fmt.Println("StartOff")
}

func (travel *TaiwanTravel) Playing() {
	fmt.Println("Playing")
}

func (travel *TaiwanTravel) Eating() {
	fmt.Println("Eating")
}

func (travel *TaiwanTravel) GoHome() {
	fmt.Println("GoHome")
}
