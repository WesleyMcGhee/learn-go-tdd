package arraysandslices

// Passes all test
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersSlice [][]int) []int {
	sumSlice := []int{}

	for _, numbers := range numbersSlice {
		sum := 0
		for _, number := range numbers {
			sum += number
		}
		sumSlice = append(sumSlice, sum)
	}

	return sumSlice
}