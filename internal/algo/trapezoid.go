package algo

import (
	"comp-math-3/internal/numeric"
	"fmt"
)

func SolveTrapezoid(ig numeric.Integral) (float64, error) {
	fmt.Printf("A: %f, B: %f, N: %d \n", ig.A, ig.B, ig.N)
	if ig.N <= 0 {
		return 0.0, fmt.Errorf("partition count must be positive")
	}

	h := (ig.B - ig.A) / float64(ig.N)

	sum := ig.F(ig.A) + ig.F(ig.B)

	for i := 1; i < ig.N; i++ {
		sum += 2 * ig.F(ig.A+float64(i)*h)
	}

	fmt.Println(sum * h / 2)
	return sum * h / 2, nil
}
