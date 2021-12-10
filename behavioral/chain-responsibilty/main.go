package main

import "pattern-golang/behavioral/chain-responsibilty/hospitals"

func main() {
	hospital := hospitals.NewHospital()
	patient := hospitals.Patient{
		Name: "ann",
	}
	hospital.MakeAnAppointment(&patient)
}
