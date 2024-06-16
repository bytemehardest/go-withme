package main

import "fmt"

func reversed(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func main() {
	A := []int{1, 2, 3, 4, 5, 6, 7}
	B := 3
	reversedNum := reversed(A)

	first := reversedNum[:B]
	second := reversedNum[B:]

	firstReverse := reversed(first)
	secondReverse := reversed(second)

	joined := append(firstReverse, secondReverse...)

	fmt.Println(joined)
}
