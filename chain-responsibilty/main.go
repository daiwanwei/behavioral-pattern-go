package main

import "behavioral-pattern-go/chain-responsibilty/hospitals"

func main() {
	hospital := hospitals.NewHospital()
	patient := hospitals.Patient{
		Name: "ann",
	}
	hospital.MakeAnAppointment(&patient)
}
