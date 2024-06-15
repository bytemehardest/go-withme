package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var pl = fmt.Println

	pl("What's your name? ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		pl("Your name is ", name)
	} else {
		log.Fatal(err)
	}
}
