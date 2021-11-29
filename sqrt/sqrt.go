package main

import (
	"fmt"
	"math"
)

func Sqrt(input float64) (float64, error) {
	if input < 0 {
		return 0.0, fmt.Errorf("ErrNegSqrt")
	} else {
		return math.Sqrt(input), nil
	}
}
