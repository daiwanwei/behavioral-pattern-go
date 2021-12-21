package main

import (
	"behavioral-pattern-go/template/travel"
	"fmt"
)

type TaiwanTravel struct {
	travel.Travel
}

func NewTaiwanTravel() *TaiwanTravel {
	t := &TaiwanTravel{}
	tmpl := travel.Travel{CustomTravel: t}
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
