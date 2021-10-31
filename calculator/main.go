package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Value 1: ")
	input1, _ := reader.ReadString('\n')
	float1 := ParseNumber(input1)

	fmt.Println("Value 2: ")
	input2, _ := reader.ReadString('\n')
	float2 := ParseNumber(input2)

	fmt.Println("Select an operation ( + - * /): ")
	sign, _ := reader.ReadString('\n')
	var result float64

	switch strings.TrimSpace(sign) {
	case "+":
		result = float1 + float2
	case "-":
		result = float1 - float2
	case "*":
		result = float1 * float2
	case "/":
		result = float1 / float2
	default:
		fmt.Println("Please provide a valid sign")
	}

	result = math.Round(result*100) / 100
	fmt.Printf("The result is %v", result)

}

func ParseNumber(numString string) float64 {
	number, err := strconv.ParseFloat(strings.TrimSpace(numString), 64)
	if err != nil {
		fmt.Println(err)
		panic("Cannot parse number")
	} else {
		return number
	}
}
