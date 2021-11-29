package main

import (
	"fmt"
)

func main() {
	var result int
	for i := 1000; i <= 9999; i++ {
		for j := 1000; j <= 9999; j++ {
			if j < i {
				continue
			}

			if isEvenEndedNumber(fmt.Sprintf("%d", i*j)) {
				result++
			}
		}
	}

	fmt.Println("The result is: ", result)
}

func isEvenEndedNumber(s string) bool {
	if s[0] == s[len(s)-1] {
		return true
	} else {
		return false
	}
}
