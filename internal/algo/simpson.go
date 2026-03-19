package algo

import (
	"comp-math-3/internal/numeric"
	"fmt"
)

func SolveSimpson(ig numeric.Integral) (float64, error) {
	if ig.N <= 0 {
		return 0.0, fmt.Errorf("partition count must be positive")
	}

	if ig.N%2 != 0 {
		return 0.0, fmt.Errorf("partition count must be even")
	}

	h := (ig.B - ig.A) / float64(ig.N)

	sum := 0.0

	for i := 0; i < ig.N-1; i += 2 {
		x0 := ig.A + float64(i)*h
		x1 := ig.A + float64(i+1)*h
		x2 := ig.A + float64(i+2)*h
		sum += ig.F(ig.A+float64(i)*h) + 4*ig.F(ig.A+float64(i+1)*h) + ig.F(ig.A+float64(i+2)*h)
		fmt.Printf("\\hline\n%d & %.3f & %.3f & %.3f & %.3f & %.3f & %.3f & %.3f \\\\ \n",
			i+1, x0, x1, x2, ig.F(x0), 4*ig.F(x1), ig.F(x2), sum*h/3)
	}

	return sum * h / 3, nil
}
