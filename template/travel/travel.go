package travel

type CustomTravel interface {
	StartOff()
	Playing()
	Eating()
	GoHome()
}

type Travel struct {
	CustomTravel
}

func (travel *Travel) Traveling() {
	travel.StartOff()
	travel.Playing()
	travel.Eating()
	travel.GoHome()
}
