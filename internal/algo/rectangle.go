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
	fmt.Printf("h = %f, x_0 = %f\n", h, startX)

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

	fmt.Println("№  x_i  y_i  s")
	for i := 0; i < n; i++ {
		x := startX + float64(i)*h
		sum += f(x)
		fmt.Printf("\\hline\n%d & %.3f & %.3f & %.3f \\\\ \n", i+1, x, f(x), sum*h)
	}

	return sum * h
}
