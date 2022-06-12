package divider

import "fmt"

func Division(num, den int) (int, error) {

	if den == 0 {
		return 0, fmt.Errorf("O denominador não pode ser 0")
	}
	result := num / den

	return result, nil

}
