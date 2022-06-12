package divider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivNumberZero(t *testing.T) {
	num1 := 5
	num2 := 0

	_, err := Division(num1, num2)
	assert.Errorf(t, err, "O denominador não pode ser 0")

}
