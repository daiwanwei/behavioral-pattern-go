package museums

import "fmt"

type WestExhibit struct {
	Name string
}

func (exhibit *WestExhibit) Welcome(visitor Visitor) {
	fmt.Println("welcome West exhibit")
	visitor.VisitForWest(exhibit)
}

func (exhibit *WestExhibit) GetName() string {
	return "West exhibit"
}
