package calc

import "errors"

// Add method returns sum of two integers
func Add(numbers ...int) (int, error) {
	sum := 0
	if len(numbers) < 2 {
		return sum, errors.New("provide more than 1 number")
	} else {
		for _, num := range numbers {
			sum = sum + num
		}
		return sum, nil
	}
}
