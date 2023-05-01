package app

import "testing"

func TestMath_Sum(t *testing.T) {
	expected := 30
	total := Sum(15, 15)

	if total != expected {
		t.Errorf("The sum result is invalid: Result: %d. Expected: %d", total, expected)
	}
}