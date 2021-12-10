package kobos

type Subscriber interface {
	NotifyNewBook(book *Book)
}

type Book struct {
	Name string
}
