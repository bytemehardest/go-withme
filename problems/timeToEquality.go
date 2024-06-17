package main

import (
	"fmt"
	"math"
)

func timeToEquality() {
	A := [...]int{2, 2, 1}
	max := math.MinInt
	sum := 0

	for _, val := range A {
		if val > max {
			max = val
		}
	}

	for _, val := range A {
		sum += max - val
	}

	fmt.Println(sum)
}
