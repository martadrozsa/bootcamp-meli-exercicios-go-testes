package ordering

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAscendingOrder(t *testing.T) {
	numbers := []int{5, 15, 11, 2, 18, 6, 9}
	result := AscendingOrder(numbers)

	expectedResult := []int{2, 5, 6, 9, 11, 15, 18}

	assert.Equal(t, expectedResult, result, "a ordem não está correta")
}
