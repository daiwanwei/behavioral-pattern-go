package devices

import "fmt"

type TV struct {
	isRunning bool
}

func (device *TV) On() {
	device.isRunning = true
	fmt.Printf("tv status %t", device.isRunning)
}

func (device *TV) Off() {
	device.isRunning = false
	fmt.Printf("tv status %t", device.isRunning)
}
