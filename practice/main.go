package main

import (
	"fmt"
)

func main() {
	anInt := 42
	p := &anInt
	fmt.Println("Value of p: ", *p)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Value 1: ", *pointer1)

	*pointer1 = 22.5
	fmt.Println("Value 1 after change: ", value1)
}
