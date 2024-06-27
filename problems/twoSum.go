package main

import (
	"fmt"
)

func twoSum() {
	nums := []int{3, 2, 4, 3, 9, 5, 1, 8}
	target := 13
	result := make([]int, 2)

	mapNums := make(map[int]int)
out:
	for index, value := range nums {
		complement := target - value

		if i, found := mapNums[complement]; found {
			result[0] = i
			result[1] = index
			break out
		}

		mapNums[value] = index
	}

	fmt.Println("two sum", result, mapNums)
}
