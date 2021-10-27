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
	fmt.Print("value 1: ")
	string1, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	value1, err := strconv.ParseFloat(strings.TrimSpace(string1), 64)

	if err != nil {
		panic("Value 1 must be a number")
	}

	fmt.Print("Value 2: ")
	string2, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	value2, err := strconv.ParseFloat(strings.TrimSpace(string2), 64)

	if err != nil {
		panic("Value 2 must be a number")
	}

	result := math.Round((value1+value2)*100) / 100

	fmt.Printf("The sum of %s and %s is %s", strconv.FormatFloat(value1, 'f', -1, 64), strconv.FormatFloat(value2, 'f', -1, 64), strconv.FormatFloat(result, 'f', -1, 64))

}
