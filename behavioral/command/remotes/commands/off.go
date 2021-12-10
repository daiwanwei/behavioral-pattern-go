package commands

import (
	"pattern-golang/behavioral/command/devices"
)

type Off struct {
	Device devices.Device
}

func (command *Off) Execute() {
	command.Device.Off()
}
