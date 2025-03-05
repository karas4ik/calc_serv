package calculators

type SimpleCalculator struct{}

func (c SimpleCalculator) Add(a, b float64) float64 {
	return a + b
}

func (c SimpleCalculator) Subtract(a, b float64) float64 {
	return a - b
}

func (c SimpleCalculator) Multiply(a, b float64) float64 {
	return a * b
}

func (c SimpleCalculator) Divide(a, b float64) float64 {
	return a / b
}
