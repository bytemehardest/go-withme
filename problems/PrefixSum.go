package main

import (
	"fmt"
)

func main() {
	A := []int{1, 2, 3, 4, 5}
	B := [][]int{{0, 3}, {1, 2}}

	prefixSum := make([]int, len(A))
	answer := make([]int, len(B))
	sum := 0

	for index, value := range A {
		sum += value
		prefixSum[index] = sum
	}

	for index, lr := range B {
		if lr[0] == 0 {
			answer[index] = prefixSum[lr[1]]
		} else {
			answer[index] = prefixSum[lr[1]] - prefixSum[lr[0]-1]
		}
	}
	fmt.Println("Prefix sum", prefixSum)
	fmt.Println("Prefix sum", answer)
}
