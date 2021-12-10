package museums

type Visitor interface {
	VisitForWest(exhibit *WestExhibit)
	VisitForEast(exhibit *EastExhibit)
}
