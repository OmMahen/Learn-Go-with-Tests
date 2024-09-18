package main

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAll(slicesOfNumber ...[]int) (result []int) {
	for _, slice := range slicesOfNumber {
		result = append(result, Sum(slice))
	}
	return
}
