package orchestrator

import (
	"fmt"
	"strconv"
	"strings"
)

type ParsedExpression struct {
	Tokens []string
}

func parseExpression(expression string) *ParsedExpression {
	tokens := strings.Fields(expression)
	return &ParsedExpression{Tokens: tokens}
}

func (p *ParsedExpression) Calculate() (float64, error) {
	if len(p.Tokens) == 0 {
		return 0, fmt.Errorf("no tokens to evaluate")
	}

	var output []float64
	var ops []string

	for i := 0; i < len(p.Tokens); i++ {
		if num, err := strconv.ParseFloat(p.Tokens[i], 64); err == nil {
			output = append(output, num)
		} else {
			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(p.Tokens[i]) {
				right := output[len(output)-1]
				output = output[:len(output)-1]
				left := output[len(output)-1]
				output = output[:len(output)-1]
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]

				result := applyOperator(left, right, op)
				output = append(output, result)
			}
			ops = append(ops, p.Tokens[i])
		}
	}

	for len(ops) > 0 {
		right := output[len(output)-1]
		output = output[:len(output)-1]
		left := output[len(output)-1]
		output = output[:len(output)-1]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		result := applyOperator(left, right, op)
		output = append(output, result)
	}

	if len(output) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return output[0], nil
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

func applyOperator(left, right float64, op string) float64 {
	switch op {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		if right == 0 {
			return 0
		}
		return left / right
	}
	return 0
}
