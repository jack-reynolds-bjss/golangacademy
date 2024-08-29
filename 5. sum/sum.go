package main

type digits struct {
	singleDigits 	[]int
	doubleDigits 	[]int
	trippleDigits 	[]int
}

func Sum(arr []int) int {
	var total int = 0;
	for _, val := range arr {
		total += val
	}
	return total
}

func SumSingleDigits(numbers digits) int {
	return Sum(numbers.singleDigits);
}

func SumDoubleeDigits(numbers digits) int {
	return Sum(numbers.doubleDigits);
}

func SumTrippleDigits(numbers digits) int {
	return Sum(numbers.trippleDigits);
}