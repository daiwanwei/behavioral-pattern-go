package logistics

import "fmt"

type Car struct {
}

func (c *Car) SendTo(to string) error {
	fmt.Printf("car send to %s", to)
	return nil
}
