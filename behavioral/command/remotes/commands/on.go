package commands

import "pattern-golang/behavioral/command/devices"

type On struct {
	Device devices.Device
}

func (command *On) Execute() {
	command.Device.On()
}
