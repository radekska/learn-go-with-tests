package arrays_slices

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersCollection ...[]int) []int {
	var sums []int
	for _, numbers := range numbersCollection {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersCollection ...[]int) []int {
	var sums []int
	for _, numbers := range numbersCollection {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}

	}
	return sums
}
