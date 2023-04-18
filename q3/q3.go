package q3

import "fmt"

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, fmt.Errorf("lista vazia")
	}
	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	numbers = append(numbers[:0], numbers[1:4]...)
	var soma float64 = 0
	for _, sum := range numbers {
		soma += float64(sum)
	}
	var media = soma / float64(len(numbers))
	return min, max, media, nil
}
