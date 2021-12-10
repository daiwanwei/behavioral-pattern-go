package hospitals

import "fmt"

type Medical struct {
	next Department
}

func (department *Medical) Execute(patient *Patient) {
	if patient.MedicineDone {
		fmt.Printf("patient %s has got medical", patient.Name)
	} else {
		fmt.Printf("patient %s get medical", patient.Name)
		patient.MedicineDone = true
	}
	department.next.Execute(patient)
}

func (department *Medical) SetNext(another Department) {
	department.next = another
}
