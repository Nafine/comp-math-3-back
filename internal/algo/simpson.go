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
		sum += ig.F(ig.A+float64(i)*h) + 4*ig.F(ig.A+float64(i+1)*h) + ig.F(ig.A+float64(i+2)*h)
	}

	return sum * h / 3, nil
}
