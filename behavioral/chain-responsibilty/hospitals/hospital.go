package hospitals

type Hospital struct {
	reception Department
	doctor    Department
	medical   Department
	cashier   Department
}

func NewHospital() *Hospital {
	reception := &Reception{}
	doctor := &Doctor{}
	medical := &Medical{}
	cashier := &Cashier{}
	reception.SetNext(doctor)
	doctor.SetNext(medical)
	medical.SetNext(cashier)
	return &Hospital{
		reception: reception,
		doctor:    doctor,
		medical:   medical,
		cashier:   cashier,
	}
}

func (h *Hospital) MakeAnAppointment(patient *Patient) {
	h.reception.Execute(patient)
}

type Department interface {
	Execute(patient *Patient)
	SetNext(another Department)
}
