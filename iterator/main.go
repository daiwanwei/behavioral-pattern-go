package main

import (
	"behavioral-pattern-go/iterator/friend"
	"fmt"
)

func main() {
	f := friend.Friend{}
	f.AddFriend(friend.User{Name: "ann"})
	f.AddFriend(friend.User{Name: "ann1"})
	f.AddFriend(friend.User{Name: "ann3"})
	iter := f.GetIterator()
	for iter.HasNext() {
		u := iter.Next()
		fmt.Println(iter.HasNext())
		fmt.Println(u)
	}
}
