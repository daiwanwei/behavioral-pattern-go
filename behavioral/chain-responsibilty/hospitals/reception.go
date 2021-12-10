package hospitals

import "fmt"

type Reception struct {
	next Department
}

func (department *Reception) Execute(patient *Patient) {
	if patient.RegistrationDone {
		fmt.Printf("patient %s registration has done", patient.Name)
	} else {
		fmt.Printf("Reception registering patient %s", patient.Name)
		patient.RegistrationDone = true
	}
	department.next.Execute(patient)
}

func (department *Reception) SetNext(another Department) {
	department.next = another
}
