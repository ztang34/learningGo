package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:])
	fmt.Println(colors)

	numbers := make([]int, 5)
	numbers[0] = 134
	numbers = append(numbers, 9)
	numbers[1] = 543

	sort.Ints(numbers)

	fmt.Println(numbers)

	var numbers2 []int
	numbers2 = append(numbers2, 5)

	fmt.Println(numbers2)

}
