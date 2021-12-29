package main

import (
	"behavioral-pattern-go/template/people"
	"behavioral-pattern-go/template/travel"
)

func main() {
	//p:= people.People{&Ann{}}
	p := people.NewPeople(&people.Ann{})
	p.Live()
	t := travel.NewTaiwanTravel()
	t.Traveling()
}
