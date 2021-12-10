package hospitals

import "fmt"

type Cashier struct {
	next Department
}

func (department *Cashier) Execute(patient *Patient) {
	if patient.PaymentDone {
		fmt.Printf("patient %s has paid", patient.Name)
	} else {
		fmt.Printf("patient %s pay", patient.Name)
		patient.PaymentDone = true
	}
	//department.next.Execute(patient)
}

func (department *Cashier) SetNext(another Department) {
	department.next = another
}
