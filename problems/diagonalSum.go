package main

import (
	"fmt"
)

func main() {

	A := [][]int{{3, 2}, {2, 3}}

	n := len(A)
	sum := 0
	for i := 0; i < n; i++ {
		sum += A[i][i]
	}

	fmt.Println("diagonal Sum", sum)
}
