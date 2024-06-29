package main

import (
	"fmt"
)

func colSum() {

	A := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 2, 3, 4}}

	m := len(A[0])
	n := len(A)
	output := make([]int, m)
	sum := 0
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			sum += A[i][j]
		}
		output[j] = sum
		sum = 0
	}

	fmt.Println("colSum", output)
}
