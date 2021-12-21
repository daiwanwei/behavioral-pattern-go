package main

import "behavioral-pattern-go/template/people"

func main() {
	//p:= people.People{&Ann{}}
	p := people.NewPeople(&Ann{})
	p.Live()
	t := NewTaiwanTravel()
	t.Traveling()
}
