package main

import (
	"fmt"
)

func main() {
	A := []int{2, 2, 2}
	B := [][]int{{0, 0}, {1, 2}}

	sumSlice := make([]int, len(B))
	for index, ranged := range B {
		numbers := A[ranged[0] : ranged[1]+1]
		sum := 0
		for _, value := range numbers {
			sum += value
		}
		sumSlice[index] = sum
	}

	fmt.Println("output: ", sumSlice)
}
