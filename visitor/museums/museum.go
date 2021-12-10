package museums

type Museum struct {
	East *EastExhibit
	West *WestExhibit
}

func NewMuseum() *Museum {
	return &Museum{
		East: &EastExhibit{},
		West: &WestExhibit{},
	}
}
