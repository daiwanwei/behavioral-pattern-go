package commands

import (
	"behavioral-pattern-go/command/devices"
)

type Off struct {
	Device devices.Device
}

func (command *Off) Execute() {
	command.Device.Off()
}
