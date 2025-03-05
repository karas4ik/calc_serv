package tests

import (
	"calc_service/internal/calculators"
	"testing"
)

func TestCalculator(t *testing.T) {
	calc := calculators.SimpleCalculator{}
	if calc.Add(2, 3) != 5 {
		t.Error("Expected 5")
	}
	if calc.Subtract(5, 3) != 2 {
		t.Error("Expected 2")
	}
	if calc.Multiply(2, 3) != 6 {
		t.Error("Expected 6")
	}
	if calc.Divide(6, 3) != 2 {
		t.Error("Expected 2")
	}
}
