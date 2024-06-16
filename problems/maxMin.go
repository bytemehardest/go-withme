package main

import (
	"fmt"
)

func maxMin() {
	A := [...]int{1, 3, 4, 1}
	max := A[0]
	min := A[0]
	for _, value := range A {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	fmt.Println(max + min)
}
