package algo

import (
	"comp-math-3/internal/numeric"
	"fmt"
)

func solveRectangleBase(ig numeric.Integral, offsetMultiplier float64) (float64, error) {
	if ig.N <= 0 {
		return 0.0, fmt.Errorf("partition count must be positive")
	}

	h := (ig.B - ig.A) / float64(ig.N)
	startX := ig.A + h*offsetMultiplier

	return accumulate(ig.F, startX, h, ig.N), nil
}

func SolveLeftRectangle(ig numeric.Integral) (float64, error) {
	return solveRectangleBase(ig, 0.0)
}

func SolveRightRectangle(ig numeric.Integral) (float64, error) {
	return solveRectangleBase(ig, 1.0)
}

func SolveMidpointRectangle(ig numeric.Integral) (float64, error) {
	return solveRectangleBase(ig, 0.5)
}

func accumulate(f func(float64) float64, startX float64, h float64, n int) float64 {
	var sum float64

	for i := 0; i < n; i++ {
		x := startX + float64(i)*h
		sum += f(x)
	}

	return sum * h
}
