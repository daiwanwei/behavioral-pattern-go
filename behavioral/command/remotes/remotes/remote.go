package remotes

import (
	"pattern-golang/behavioral/command/devices"
	"pattern-golang/behavioral/command/remotes/buttons"
	"pattern-golang/behavioral/command/remotes/commands"
)

type Remote struct {
	OnButton  buttons.Button
	OffButton buttons.Button
}

func NewRemote() *Remote {
	device := &devices.TV{}
	on := buttons.Button{
		Command: &commands.On{Device: device},
	}
	off := buttons.Button{
		Command: &commands.Off{Device: device},
	}
	return &Remote{
		on,
		off,
	}
}
