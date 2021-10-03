package array

func Sum(numbers []int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}

func _SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	length := len(numbersToSum)
	sums = make([]int, length)
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}
