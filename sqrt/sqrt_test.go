package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return math.Abs(v1-v2) <= 0.001
}

type testCases struct {
	input          float64
	expectedOutput float64
}

func getTestCases(filename string) *[]testCases {
	file, _ := os.Open(filename)
	defer file.Close()

	result := []testCases{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		input, _ := strconv.ParseFloat(line[0], 64)
		expected, _ := strconv.ParseFloat(line[1], 64)
		t := testCases{input, expected}
		result = append(result, t)
	}

	return &result
}

func Test(t *testing.T) {
	cases := getTestCases("sqrt_cases.csv")
	for _, tc := range *cases {
		t.Run(fmt.Sprintf("%f", tc.input), func(t *testing.T) {
			out, err := Sqrt(tc.input)
			if err != nil {
				t.Fatal("error")
			}

			if !almostEqual(out, tc.expectedOutput) {
				t.Fatalf("%f != %f", out, tc.expectedOutput)
			}
		})
	}
}

func BenchMarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Sqrt(float64(i))

		if err != nil {
			b.Fatal(err)
		}
	}
}
