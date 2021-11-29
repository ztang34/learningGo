package main

import (
	"fmt"
)

func main() {

	/**Introduction
	fmt.Println("Welcome Gophers :)")

	var x int
	var y int
	x = 3
	y = 4

	fmt.Printf("X is %v, y is %v, x is of type %T, y is of type %T ", x, y, x, y) **/

	/**conditional statement
	x := 5
	y := 2
	mean := (float64(x) + float64(y)) / 2.0

	fmt.Printf("Mean is %v, type of mean is %T", mean, mean)

	if mean > 5 {
		fmt.Println("Mean is greater than 5")
	} else {
		fmt.Println("Mean is less than 5")
	}

	if sum := x + y; sum > 5 {
		fmt.Println("Sum is gretan than 5")
	} else {
		fmt.Println("SUm is less than 5")
	}

	switch {
	case mean > 5:
		fmt.Println("Mean is greater than 5")
	default:
		fmt.Println("Mean is less than 5")
	}**/

	/** loop
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	i := 0

	for i < 3 {
		fmt.Println(i)
		i++
	}

	i = 0

	for {
		if i > 3 {
			break
		} else {
			fmt.Println(i)
		}

		i++
	}**/

	/** string practice
	book := "The color of magic"
	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Printf("The third character is %v, its type is %T", book[2], book[2])
	fmt.Printf("The third character is %c, its type is %T", book[2], book[2])
	fmt.Println(book[:len(book)-1])

	for _, i := range book {
		fmt.Println(string(i))
	}

	mutliline := `This is
	how you can use
	back tick to assign mulltipe line string`
	fmt.Println(mutliline)**/

	/** slice
	slice := []string{"abc", "kid", "tv"}
	fmt.Printf("The value of slice is %v, the type is %T\n", slice, slice)

	slice = append(slice, "wee!")

	fmt.Println(slice[1:])

	for _, i := range slice {
		fmt.Println(i)
	}**/

	stocks := map[string]float64{
		"AMZN": 1,
		"MSFT": 2,
		"AAPL": 3,
	}

	stocks["AMZN"] = 25
	stocks["GOOG"] = 250
	delete(stocks, "AMZN")
	delete(stocks, "TSLA")

	fmt.Println(len(stocks))

	for key, value := range stocks {
		fmt.Printf("The price for %s is %v\n", key, value)
	}

	price, ok := stocks["TSLA"]

	if ok {
		fmt.Println(price)
	} else {
		fmt.Println("No price for TSLA!")
	}

	stocks["TSLA"] = 100
	price, ok = stocks["TSLA"]

	if ok {
		fmt.Println(price)
	} else {
		fmt.Println("No price for TSLA!")
	}

}
