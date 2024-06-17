package main

import (
	"fmt"
	"math"
)

func secondLargestNumber() {
	A := [...]int{13, 7, 16, 18, 14, 17, 18, 8, 10}
	largest := A[0]
	secondLargest := math.MinInt
	if len(A) == 0 {
		fmt.Println(-1)
	}

	for _, value := range A {
		if value > largest {
			secondLargest = largest
			largest = value
		}
		if value < largest && value > secondLargest {
			secondLargest = value
		}
	}

	if secondLargest == math.MinInt {
		fmt.Println(-1)
	}

	fmt.Println(secondLargest)
}
