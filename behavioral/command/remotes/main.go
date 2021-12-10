package main

import "pattern-golang/behavioral/command/remotes/remotes"

func main() {
	remote := remotes.NewRemote()
	remote.OnButton.Press()
	remote.OffButton.Press()
}
