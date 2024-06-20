package main

import (
	"fmt"
)

func main() {
	A := []int{1, 2, 3, 4, 5}
	B := [][]int{{0, 2}, {2, 4}, {1, 4}}

	prefixSum := make([]int, len(A))
	evens := make([]int, len(B))

	if A[0]%2 == 1 {
		prefixSum[0] = 0
	} else {
		prefixSum[0] = 1
	}

	for i := 1; i < len(A); i++ {
		if A[i]%2 == 1 {
			prefixSum[i] = prefixSum[i-1]
		} else {
			prefixSum[i] = prefixSum[i-1] + 1
		}
	}

	for index, value := range B {
		evens[index] = prefixSum[value[1]] - prefixSum[value[0]]
	}

	fmt.Println(evens)
}
