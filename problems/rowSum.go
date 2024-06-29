package main

import (
	"fmt"
)

func rowSum() {

	A := [][]int{{1, -2, -3}, {-4, 5, -6}, {-7, -8, 9}}

	m := len(A[0])
	n := len(A)
	output := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sum += A[i][j]
		}
		output[i] = sum
		sum = 0
	}

	fmt.Println("rowSum", output)
}
