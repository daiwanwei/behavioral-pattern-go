package museums

import "fmt"

type EastExhibit struct {
}

func (exhibit *EastExhibit) Welcome(visitor Visitor) {
	fmt.Println("welcome East exhibit")
	visitor.VisitForEast(exhibit)
}

func (exhibit *EastExhibit) GetName() string {
	return "East exhibit"
}
