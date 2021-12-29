package travel

type customTravel interface {
	StartOff()
	Playing()
	Eating()
	GoHome()
}

type Travel struct {
	custom customTravel
}

func (travel *Travel) Traveling() {
	travel.custom.StartOff()
	travel.custom.Playing()
	travel.custom.Eating()
	travel.custom.GoHome()
}
