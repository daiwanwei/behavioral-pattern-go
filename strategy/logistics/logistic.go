package logistics

type Logistic interface {
	SendTo(to string) error
}
