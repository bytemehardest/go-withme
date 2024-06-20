package main

import (
	"fmt"
)

func evenInRange() {
	A := []int{2, 1, 8, 3, 9, 6}
	B := [][]int{{0, 3}, {3, 5}, {1, 3}, {2, 4}}

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

	fmt.Println(prefixSum)

	for index, value := range B {
		if value[0] == 0 {
			evens[index] = prefixSum[value[1]]
		} else {
			evens[index] = prefixSum[value[1]] - prefixSum[value[0]-1]
		}
	}

	fmt.Println(evens)
}
