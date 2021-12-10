package kobos

import "fmt"

type User struct {
	Name string
}

func (u *User) NotifyNewBook(book *Book) {
	fmt.Printf("%s receive notification of book(%s)\n", u.Name, book.Name)
}
