package kobos

type BookStore struct {
	Subscribers []Subscriber
}

func (store *BookStore) Subscribe(subscriber Subscriber) {
	store.Subscribers = append(store.Subscribers, subscriber)
}

func (store *BookStore) Unsubscribe(subscriber Subscriber) {
	var index int
	for i, s := range store.Subscribers {
		if s == subscriber {
			index = i
			break
		}
	}
	store.Subscribers = append(store.Subscribers[:index], store.Subscribers[index+1:]...)
}

func (store *BookStore) Notify(book *Book) {
	for _, s := range store.Subscribers {
		s.NotifyNewBook(book)
	}
}
