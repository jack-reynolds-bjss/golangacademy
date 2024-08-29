package main

import (
	"sort"
)

type OddEven string;

const (
	odd OddEven = "odd"
	even OddEven = "even"
	both OddEven = "both"
)

func IsOddEven(val int, oddEven OddEven) bool {
	return oddEven == both || (oddEven == even && val % 2 == 0) || (oddEven == odd && val % 2 == 1)
}

func FilterArray(arr []int, oddEven OddEven) (result []int) {
	for _, val := range arr {
		if IsOddEven(val, oddEven) {
			result = append(result, val)
		}
	}
	return
}

func FilterAndSortAscending(arr []int, oddEven OddEven) (result []int) {
	result = FilterArray(arr, oddEven)
	sort.Slice(result, func(i int, j int) bool {
		return result[i] < result[j]
	})
	return
}

func FilterAndSortDescending(arr []int, oddEven OddEven) (result []int) {
	result = FilterArray(arr, oddEven)
	sort.Slice(result, func(i int, j int) bool {
		return result[i] > result[j]
	})
	return
}