package people

import "fmt"

type People struct {
	customPeople
}

func (p *People) Live() {
	p.WakeUp()
	p.Eat()
	p.Working()
	p.Eat()
	p.Working()
	p.GoHome()
	p.Eat()
	p.Hangout()
	p.Sleep()
}

func NewPeople(people customPeople) *People {
	return &People{people}
}

type customPeople interface {
	WakeUp()
	Eat()
	Working()
	GoHome()
	Hangout()
	Sleep()
}

type templatePeople struct {
}

func (p *templatePeople) Hangout() {
	fmt.Printf("hangout\n")
}
