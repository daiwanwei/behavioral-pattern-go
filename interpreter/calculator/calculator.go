package calculator

import (
	"strconv"
	"strings"
)

type Calculator struct {
}

func (c *Calculator) Calculate(expression string) int {
	stack := expressionStack{}
	es := strings.Split(expression, " ")
	for _, e := range es {
		switch e {
		case "+":
			right := stack.Pop()
			left := stack.Pop()
			sum := Sum{left, right}
			value := Value(sum.Calculate())
			stack.Push(&value)
		case "-":
			right := stack.Pop()
			left := stack.Pop()
			sum := Subtract{left, right}
			value := Value(sum.Calculate())
			stack.Push(&value)
		default:
			val, err := strconv.Atoi(e)
			if err != nil {
				panic(err)
			}
			v := Value(val)
			stack.Push(&v)
		}
	}
	return stack.Pop().Calculate()
}

type expressionStack []Expression

func (stack *expressionStack) Pop() Expression {
	l := len(*stack)
	if l <= 0 {
		return nil
	}
	e := (*stack)[l-1]
	*stack = (*stack)[:l-1]
	return e
}

func (stack *expressionStack) Push(e Expression) {
	*stack = append(*stack, e)
}
