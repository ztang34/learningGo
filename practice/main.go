package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")
	doSomething()
	sum := addValues(5, 8)
	fmt.Println(sum)

	result, number := addAllValues(1, 2, 3, 4, 5)

	fmt.Println(result, number)

}

func doSomething() {
	fmt.Println("Doing something")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, i := range values {
		total += i
	}

	return total, len(values)
}
