package main

import (
	"fmt"
	"strings"
)

func main() {
	var pl = fmt.Println

	/* pl("What's your name? ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		pl("Your name is ", name)
	} else {
		log.Fatal(err)
	} */

	// type of
	// pl(reflect.TypeOf(25))

	// Conditionals
	/*
		var age int = 24
		if age > 1 && age <= 18 {
			pl("Important birthday")
		} else if age == 21 || age == 50 {
			pl("Important birthday")
		} else if age >= 65 {
			pl("Important birthday")
		} else {
			pl("Not Important birthday")
		}
	*/

	//Strings
	sV1 := "A string"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)
	pl("Length: ", len(sV2))
	pl("Contains Another: ", strings.Contains(sV2, "Another"))
	pl("o index: ", strings.Index(sV2, "o"))
	pl("o Replace: ", strings.Replace(sV2, "o", "0", -1))
	sV3 := "1-2-3-4-5"
	pl("Split", strings.Split(sV3, "-"))
	pl("Prefix", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix", strings.HasSuffix("tacocat", "cat"))

}
