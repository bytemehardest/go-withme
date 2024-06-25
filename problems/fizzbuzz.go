package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := 50
	resultSlice := make([]string, input)

	for i := 1; i <= input; i++ {
		if i%15 == 0 {
			resultSlice = append(resultSlice, "FizzBuzz")
		} else if i%3 == 0 {
			resultSlice = append(resultSlice, "Fizz")
		} else if i%5 == 0 {
			resultSlice = append(resultSlice, "Buzz")
		} else {
			resultSlice = append(resultSlice, strconv.Itoa(i))
		}
	}

	fmt.Println("fizz buzz", resultSlice)
}
