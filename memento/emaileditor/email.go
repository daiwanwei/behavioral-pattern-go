package emaileditor

type Email struct {
	state Content
}

func (memento Email) GetState() Content {
	return memento.state
}

type CurrentEmail struct {
	Email
}

func (originator *CurrentEmail) save() Email {
	return originator.Email
}

func (originator *CurrentEmail) load(email Email) {
	originator.Email = email
}

func (originator *CurrentEmail) SetTitle(title string) {
	originator.state.Title = title
}

func (originator *CurrentEmail) SetBody(body string) {
	originator.state.Body = body
}

type Content struct {
	Title string
	Body  string
}
