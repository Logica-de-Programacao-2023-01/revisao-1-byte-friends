package q3

import "fmt"

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, fmt.Errorf("lista vazia")
	}
	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num < max {
			max = num
		}
		if num > min {
			min = num
		}
	}
	var soma float64 = 0

	for _, sum := range numbers {
		soma += float64(sum)
	}
	return min, max, (soma - (float64(min)) - float64(max)) / float64(len(numbers)-2), nil
}
