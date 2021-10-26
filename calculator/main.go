package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("value 1: ")
	string1, _ := bufio.NewReader(os.Stdin).ReadString('\r')
	value1, err := strconv.ParseInt(string1, 10, 64)

	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	fmt.Print("Value 2: ")
	string2, _ := bufio.NewReader(os.Stdin).ReadString('\r')
	value2, err := strconv.ParseFloat(string2, 64)

	if err != nil {
		panic("Value 2 must be a number")
	}

	fmt.Println(string1, string2, value1, value2)

}
