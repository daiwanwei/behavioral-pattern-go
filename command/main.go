package main

import (
	"behavioral-pattern-go/command/remotes"
)

func main() {
	remote := remotes.NewRemote()
	remote.OnButton.Press()
	remote.OffButton.Press()
}
