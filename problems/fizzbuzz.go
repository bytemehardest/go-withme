package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 50
	resultSlice := make([]string, n)

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			resultSlice[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			resultSlice[i-1] = "Fizz"
		} else if i%5 == 0 {
			resultSlice[i-1] = "Buzz"
		} else {
			resultSlice[i-1] = strconv.Itoa(i)
		}
	}

	fmt.Println("fizz buzz", resultSlice)
}
