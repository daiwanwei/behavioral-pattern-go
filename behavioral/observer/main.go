package main

import "pattern-golang/behavioral/observer/kobos"

func main() {
	book := kobos.Book{
		Name: "dog vs cat",
	}
	store := kobos.BookStore{}
	store.Notify(&book)

	user := kobos.User{
		Name: "ann",
	}

	writer := kobos.Writer{
		Name: "ann",
	}
	store.Subscribe(&user)
	store.Subscribe(&writer)
	store.Notify(&book)
}
