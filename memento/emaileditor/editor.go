package emaileditor

type Editor struct {
	Current   CurrentEmail
	histories []Email
}

func NewEditor() *Editor {
	return &Editor{
		Current:   CurrentEmail{Email{state: Content{"", ""}}},
		histories: make([]Email, 0),
	}
}

func (caretaker *Editor) Save() {
	current := caretaker.Current.save()
	caretaker.histories = append(caretaker.histories, current)
}

func (caretaker *Editor) Undo() {
	if len(caretaker.histories) == 0 {
		return
	}
	email := caretaker.histories[len(caretaker.histories)-1]
	caretaker.histories = caretaker.histories[:len(caretaker.histories)-1]
	caretaker.Current = CurrentEmail{email}
}
