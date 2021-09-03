package main


func Sum(numbers []int) int {
	sum := 0

	for i, _ := range(numbers) {
		sum += numbers[i]
	}

    return sum
}


func SumAll(numbersToSum ...[]int) []int {
    lengthOfNumbers := len(numbersToSum)
    sums := make([]int, lengthOfNumbers)

    for i, numbers := range numbersToSum {
        sums[i] = Sum(numbers)
    }

    return sums
}
