package main

import (
	"fmt"
)

func reversed(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func rightRotate() {
	A := []int{1, 1, 4, 9, 4, 7, 1}
	B := 9
	rotations := B % len(A)

	reversedNum := reversed(A)

	firstReverse := reversed(reversedNum[:rotations])
	secondReverse := reversed(reversedNum[rotations:])

	joined := append(firstReverse, secondReverse...)

	fmt.Println(joined)
}
