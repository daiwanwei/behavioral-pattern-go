package museums

type Exhibit interface {
	Welcome(visitor Visitor)
	GetName() string
}
