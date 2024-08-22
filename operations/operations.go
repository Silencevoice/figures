package operations

import "fmt"

// Operation models an operation with two operands, and operator and a result.
type Operation struct {
	FirstOperand  int
	SecondOperand int
	Operator      string
	Result        int
}

// String method for Operation
func (op Operation) String() string {
	return fmt.Sprintf("%d %s %d = %d", op.FirstOperand, op.Operator, op.SecondOperand, op.Result)
}

// Sum performs addition between two integers and returns an Operation.
func Sum(a, b int) Operation {
	return Operation{FirstOperand: a, SecondOperand: b, Operator: "+", Result: a + b}
}

// Subtract performs subtraction between two integers and returns an Operation if valid.
func Subtract(a, b int) (Operation, bool) {
	if a >= b {
		return Operation{FirstOperand: a, SecondOperand: b, Operator: "-", Result: a - b}, true
	}
	return Operation{}, false
}

// Multiply performs multiplication between two integers and returns an Operation.
func Multiply(a, b int) Operation {
	return Operation{FirstOperand: a, SecondOperand: b, Operator: "*", Result: a * b}
}

// Divide performs division between two integers and returns an Operation if valid.
func Divide(a, b int) (Operation, bool) {
	if b != 0 && a%b == 0 {
		return Operation{FirstOperand: a, SecondOperand: b, Operator: "/", Result: a / b}, true
	}
	return Operation{}, false
}
