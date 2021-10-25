package main

import (
	"fmt"
	"math"
)

func main() {
	i1, i2, i3 := 12, 45, 68

	intSum := i1 + i2 + i3

	fmt.Println("Integer sum is: ", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("float sum is: ", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("Rounded float sum is: ", floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference)

}
