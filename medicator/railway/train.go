package railway

import (
	"fmt"
)

type Train interface {
	Arrive()
	Depart()
	NotifyAboutArrival()
}

type PassengerTrain struct {
	Mediator Mediator
}

func (train *PassengerTrain) Arrive() {
	if train.Mediator.CanArrival(train) {
		fmt.Println("PassengerTrain: Arrived")
		return
	}
	fmt.Println("PassengerTrain: Arrival blocked, waiting")
}

func (train *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	train.Mediator.NotifyAboutDeparture()
}

func (train *PassengerTrain) NotifyAboutArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	train.Arrive()
}

type FreightTrain struct {
	Mediator Mediator
}

func (train *FreightTrain) Arrive() {
	if train.Mediator.CanArrival(train) {
		fmt.Println("FreightTrain: Arrived")
		return
	}
	fmt.Println("FreightTrain: Arrival blocked, waiting")
}

func (train *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	train.Mediator.NotifyAboutDeparture()
}

func (train *FreightTrain) NotifyAboutArrival() {
	fmt.Println("FreightTrain: Arrival permitted, arriving")
	train.Arrive()
}
