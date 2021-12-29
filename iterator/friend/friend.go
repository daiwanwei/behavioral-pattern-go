package friend

type Friend struct {
	users []User
}

func (collection *Friend) AddFriend(user User) {
	collection.users = append(collection.users, user)
}

func (collection *Friend) GetIterator() Iterator {
	return &UserIterator{index: 0, users: collection.users}
}

type UserIterator struct {
	index int
	users []User
}

func (iterator *UserIterator) HasNext() bool {

	return iterator.index < len(iterator.users)
}

func (iterator *UserIterator) Next() *User {
	if iterator.HasNext() {
		u := iterator.users[iterator.index]
		iterator.index++
		return &u
	}
	return nil
}

type Collection interface {
	GetIterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() *User
}
