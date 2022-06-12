package ordering

import "sort"

func AscendingOrder(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}
