package kobos

import "fmt"

type Writer struct {
	Name string
}

func (w *Writer) NotifyNewBook(book *Book) {
	fmt.Printf("%s receive notification of book(%s)\n", w.Name, book.Name)
}
