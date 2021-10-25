package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()

	fmt.Println("The current time is: ", n)

	t := time.Date(2021, time.November, 10, 24, 0, 0, 0, time.UTC)

	fmt.Println(t)
	fmt.Println(t.Format(time.ANSIC))

	t2, _ := time.Parse(time.ANSIC, "Thu Nov 11 00:00:00 2021")
	fmt.Printf("The type is %T", t2)

	fmt.Println("Dates and times")

}
