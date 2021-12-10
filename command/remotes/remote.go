package remotes

import (
	"behavioral-pattern-go/command/devices"
	"behavioral-pattern-go/command/remotes/buttons"
	"behavioral-pattern-go/command/remotes/commands"
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
