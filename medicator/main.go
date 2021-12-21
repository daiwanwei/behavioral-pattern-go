package main

import (
	"behavioral-pattern-go/medicator/railway"
)

func main() {
	manger := railway.NewStationManager()
	train1 := railway.PassengerTrain{Mediator: manger}
	train2 := railway.FreightTrain{Mediator: manger}
	train1.Arrive()
	train2.Arrive()
	train1.Depart()
	train2.Depart()
}
