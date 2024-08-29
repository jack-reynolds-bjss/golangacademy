package main

import (
	"reflect"
	"testing"
)

type OddEvenExample struct {
	Value		int
	OddEven 	OddEven
	Expected 	bool
}

var OddEvenExamples = []OddEvenExample{
	{0, odd, false},
	{0, even, true},
	{0, both, true},
	{1, odd, true},
	{1, even, false},
	{1, both, true},
	{2, odd, false},
	{2, even, true},
	{2, both, true},
	{3, odd, true},
	{3, even, false},
	{3, both, true},
}

type FilterArrayExample struct {
	Value		[]int
	OddEven 	OddEven
	Expected 	[]int
}

var FilterArrayExamples = []FilterArrayExample{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, odd, []int{1, 3, 5, 7, 9}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, even, []int{2, 4, 6, 8, 10}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, both, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
}

type FilterAndSortArrayAscendingExample struct {
	Value		[]int
	OddEven 	OddEven
	Expected 	[]int
}

var FilterAndSortArrayAscendingExamples = []FilterAndSortArrayAscendingExample{
	{[]int{4, 7, 2, 3, 1, 10, 8, 5, 9, 6}, odd, []int{1, 3, 5, 7, 9}},
	{[]int{4, 7, 2, 3, 1, 10, 8, 5, 9, 6}, even, []int{2, 4, 6, 8, 10}},
	{[]int{4, 7, 2, 3, 1, 10, 8, 5, 9, 6}, both, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
}

type FilterAndSortArrayDescendingExample struct {
	Value		[]int
	OddEven 	OddEven
	Expected 	[]int
}

var FilterAndSortArrayDescendingExamples = []FilterAndSortArrayDescendingExample{
	{[]int{4, 7, 2, 3, 1, 10, 8, 5, 9, 6}, odd, []int{9, 7, 5, 3, 1}},
	{[]int{4, 7, 2, 3, 1, 10, 8, 5, 9, 6}, even, []int{10, 8, 6, 4, 2}},
	{[]int{4, 7, 2, 3, 1, 10, 8, 5, 9, 6}, both, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
}

func TestFilterAndSortArray(t *testing.T) {
	t.Run("IsOddEven works correctly", func(t *testing.T) {
		for _, example := range OddEvenExamples {
			result := IsOddEven(example.Value, example.OddEven)
			if result != example.Expected {
				t.Error(
					"For", example.Value,
					"and", example.OddEven,
					"Expected", example.Expected,
					"Got", result,
				)
			}
		}
	})

	t.Run("FilterArray works correctly", func(t *testing.T) {
		for _, example := range FilterArrayExamples {
			result := FilterArray(example.Value, example.OddEven)
			if !reflect.DeepEqual(result, example.Expected) {
				t.Error(
					"For", example.Value,
					"and", example.OddEven,
					"Expected", example.Expected,
					"Got", result,
				)
			}
		}
	})



	t.Run("FilterAndSortAscending works correctly", func(t *testing.T) {
		for _, example := range FilterAndSortArrayAscendingExamples {
			result := FilterAndSortAscending(example.Value, example.OddEven)
			if !reflect.DeepEqual(result, example.Expected) {
				t.Error(
					"For", example.Value,
					"and", example.OddEven,
					"Expected", example.Expected,
					"Got", result,
				)
			}
		}
	})

	t.Run("FilterAndSortDescending works correctly", func(t *testing.T) {
		for _, example := range FilterAndSortArrayDescendingExamples {
			result := FilterAndSortDescending(example.Value, example.OddEven)
			if !reflect.DeepEqual(result, example.Expected) {
				t.Error(
					"For", example.Value,
					"and", example.OddEven,
					"Expected", example.Expected,
					"Got", result,
				)
			}
		}
	})
}