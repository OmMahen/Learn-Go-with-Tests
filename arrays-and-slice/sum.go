package main

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAllTails(slicesOfNumber ...[]int) (result []int) {
	for _, slice := range slicesOfNumber {
		if len(slice) == 0 {
			result = append(result, 0)

		} else {
			tail := slice[1:]
			result = append(result, Sum(tail))
		}
	}
	return
}
