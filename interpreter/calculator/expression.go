package calculator

type Expression interface {
	Calculate() int
}

type Value int

func (expression *Value) Calculate() int {
	return int(*expression)
}

type Sum struct {
	Left  Expression
	Right Expression
}

func (expression *Sum) Calculate() int {
	return expression.Left.Calculate() + expression.Right.Calculate()
}

type Subtract struct {
	Left  Expression
	Right Expression
}

func (expression *Subtract) Calculate() int {
	return expression.Left.Calculate() - expression.Right.Calculate()
}
