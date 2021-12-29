package people

import (
	"fmt"
)

type Ann struct {
	templatePeople
}

func (people *Ann) WakeUp() {
	fmt.Printf("%s wake up at 6:00\n", "ann")
}

func (people *Ann) Eat() {
	fmt.Printf("%s eat \n", "ann")
}

func (people *Ann) Working() {
	fmt.Printf("%s work\n", "ann")
}

func (people *Ann) GoHome() {
	fmt.Printf("%s go home at 18:00\n", "ann")
}

//func (people *Ann) Hangout() {
//	fmt.Printf("%s hangout\n","ann")
//}

func (people *Ann) Sleep() {
	fmt.Printf("%s sleep at 22:00\n", "ann")
}
