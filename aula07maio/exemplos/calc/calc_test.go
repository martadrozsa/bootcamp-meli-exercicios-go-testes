package calc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumPositiveValues(t *testing.T) {
	num1 := 3
	num2 := 5
	expectedResult := 8

	result := Sum(num1, num2)

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestSumNegativeValues(t *testing.T) {
	num1 := -3
	num2 := -5
	expectedResult := -8

	result := Sum(num1, num2)

	if result != expectedResult {
		t.Errorf("A função Sum() retornou o resultado = %v, mas o esperado é %v", result, expectedResult)
	}
}

func TestSumPositiveAndNegativeValues(t *testing.T) {
	num1 := 3
	num2 := -5
	expectedResult := -2

	result := Sum(num1, num2)

	if result != expectedResult {
		t.Errorf("A função Sum() retornou o resultado = %v, mas o esperado é %v", result, expectedResult)
	}
}

func TestSumNegativeAndPositiveValues(t *testing.T) {
	num1 := -3
	num2 := 5
	expectedResult := 2

	result := Sum(num1, num2)

	if result != expectedResult {
		t.Errorf("A função Sum() retornou o resultado = %v, mas o esperado é %v", result, expectedResult)
	}
}

// TODO
/*
func TestSumPositiveMaxValues(t *testing.T) {
	num1 := 9999
	num2 := 9999
	expectedResult := 2

	result := Sum(num1, num2)

	if result != expectedResult {
		t.Errorf("A função Sum() retornou o resultado = %v, mas o esperado é %v", result, expectedResult)
	}
}
*/

func TestSub(t *testing.T) {
	num1 := 15
	num2 := 5
	expectedResult := 10

	result := Sub(num1, num2)

	if result != expectedResult {
		t.Errorf("A função Sub() retornou o resultado = %v, mas o esperado é %v", result, expectedResult)
	}
}
