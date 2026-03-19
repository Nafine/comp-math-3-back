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

	sum := ig.F(ig.A)

	for i := 1; i < ig.N; i++ {
		sum += 2 * ig.F(ig.A+float64(i)*h)
		fmt.Printf("\\hline\n%d & %.3f & %.3f & %.3f & %.3f & %.3f  \\\\ \n",
			i, ig.A+float64(i-1)*h, ig.A+float64(i)*h, ig.F(ig.A+float64(i-1)*h), ig.F(ig.A+float64(i)*h), sum*h/2)
	}

	sum += ig.F(ig.B)

	fmt.Printf("\\hline\n%d & %.3f & %.3f & %.3f & %.3f & %.3f  \\\\ \n",
		ig.N, ig.A+float64(ig.N-1)*h, ig.A+float64(ig.N)*h, ig.F(ig.A+float64(ig.N-1)*h), ig.F(ig.A+float64(ig.N)*h), sum*h/2)

	fmt.Println(sum * h / 2)
	return sum * h / 2, nil
}
