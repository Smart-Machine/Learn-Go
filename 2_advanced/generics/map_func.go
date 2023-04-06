package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// input:  [1, 2, 3]; (n) => n * 2
// output: [2, 4, 6]

func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {
	result := MapValues([]int{1, 2, 3}, func(n int) int {
		return n * 5
	})
	fmt.Printf("result: %+v\n", result)
}
