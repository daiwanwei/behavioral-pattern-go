package railway

type Mediator interface {
	CanArrival(t Train) bool
	NotifyAboutDeparture()
}

type StationManager struct {
	trains         []Train
	isPlatformFree bool
}

func NewStationManager() Mediator {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (mediator *StationManager) CanArrival(t Train) bool {
	if mediator.isPlatformFree {
		mediator.isPlatformFree = false
		return true
	}
	mediator.trains = append(mediator.trains, t)
	return false
}

func (mediator *StationManager) NotifyAboutDeparture() {
	if !mediator.isPlatformFree {
		mediator.isPlatformFree = false
	}
	if len(mediator.trains) > 0 {
		next := mediator.trains[0]
		mediator.trains = mediator.trains[1:]
		next.NotifyAboutArrival()
	}
}
