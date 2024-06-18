package main

import (
	"fmt"
)

func prefixSum() {
	A := []int{1, 2, 3, 4, 5}
	B := [][]int{{0, 3}, {1, 2}}

	prefixSum := make([]int, len(A))
	answer := make([]int64, len(B))
	sum := 0

	for index, value := range A {
		sum += value
		prefixSum[index] = sum
	}

	for index, lr := range B {
		if lr[0] == 0 {
			answer[index] = int64(prefixSum[lr[1]])
		} else {
			answer[index] = int64(prefixSum[lr[1]]) - int64(prefixSum[lr[0]-1])
		}
	}
	fmt.Println("Prefix sum", prefixSum)
	fmt.Println("Prefix sum", answer)
}
