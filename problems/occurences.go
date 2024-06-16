package main

import (
	"fmt"
)

func occurences() {
	A := [...]int{1, 3, 4, 1, 2, 2, 5, 5, 5, 5, 5, 5, 5, 6}
	B := 5
	occurences := 0

	for _, value := range A {
		if value == B {
			occurences++
		}
	}

	fmt.Println(occurences)
}
