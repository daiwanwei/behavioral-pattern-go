package commands

import "behavioral-pattern-go/command/devices"

type On struct {
	Device devices.Device
}

func (command *On) Execute() {
	command.Device.On()
}
