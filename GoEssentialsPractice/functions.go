package main

import (
	"fmt"
	"math"
)

func main() {
	a := 24
	b := 9

	fmt.Println(sum(a, b))
	fmt.Println(divide(a, b))

	nums := []int{1, 3, 5, 7, 9, 11}
	doubleAt(nums, 3)
	fmt.Println(nums)
	num := nums[3]
	double(&num)
	fmt.Println(num)

	result, err := sqrt(5.5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func sum(a, b int) int {
	return a + b
}

func divide(a, b int) (int, int) {
	return a / b, a % b

}

func doubleAt(nums []int, i int) {
	nums[i] *= 2
}

func double(n *int) {
	*n *= 2
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("Cannot calculate sqrt for negative value")
	} else {
		return math.Sqrt(n), nil
	}
}
