package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional logic")

	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "less than zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	if x := -42; x < 0 {
		result = "Less than zero"
	} else {
		result = "greater or equal to zero"
	}

	fmt.Println(result)
}
