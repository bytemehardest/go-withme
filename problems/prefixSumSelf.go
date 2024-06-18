package main

import (
	"fmt"
)

func prefixSumSelf() {

	A := []int{4, 3, 2}

	for index := range A {
		if index != 0 {
			A[index] += A[index-1]
		}
	}

	fmt.Println("Prefix Sum self: ", A)
}
