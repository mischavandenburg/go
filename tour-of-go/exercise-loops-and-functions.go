package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	n := 40.0
	fmt.Println("Result from Sqrt is ", Sqrt(n))
	fmt.Println("Result from math.Sqrt is ", math.Sqrt(n))
}
