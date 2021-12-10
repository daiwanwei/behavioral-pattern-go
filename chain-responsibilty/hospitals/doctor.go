package hospitals

import "fmt"

type Doctor struct {
	next Department
}

func (department *Doctor) Execute(patient *Patient) {
	if patient.DoctorCheckUpDone {
		fmt.Printf("patient %s doctor check up done", patient.Name)
	} else {
		fmt.Printf("doctor check up patient %s", patient.Name)
		patient.DoctorCheckUpDone = true
	}
	department.next.Execute(patient)
}

func (department *Doctor) SetNext(another Department) {
	department.next = another
}
