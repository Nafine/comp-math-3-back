package main

import (
	"comp-math-3/internal/algo"
	"comp-math-3/internal/numeric"
	"fmt"
)

func main() {
	ig := numeric.Integral{
		F: func(x float64) float64 {
			return x * x
		},
		Tolerance: 0.000001,
		N:         4,
		A:         1,
		B:         2,
	}

	fmt.Println(algo.Solve("rectangleRight", ig))
	fmt.Println(algo.Solve("rectangleLeft", ig))
	fmt.Println(algo.Solve("rectangleMidpoint", ig))
	fmt.Println(algo.Solve("trapezoid", ig))
	fmt.Println(algo.Solve("simpson", ig))
}
