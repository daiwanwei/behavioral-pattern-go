package logistics

import "fmt"

type Bike struct {
}

func (c *Bike) SendTo(to string) error {
	fmt.Printf("bike send to %s", to)
	return nil
}
