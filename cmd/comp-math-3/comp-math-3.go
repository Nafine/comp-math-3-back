package main

import (
	"comp-math-3/internal/algo"
	"comp-math-3/internal/numeric"
	"fmt"
)

func main() {
	//cfg, err := config.Get()
	//
	//if err != nil {
	//	panic(err)
	//}

	ig := numeric.Integral{
		F:         numeric.GetFunction(4),
		Tolerance: 0.001,
		N:         4,
		A:         -1,
		B:         3,
	}

	fmt.Println(algo.Solve("trapezoid", ig))

	ig2 := numeric.Integral{
		F:         numeric.GetFunction(2),
		Tolerance: 0.001,
		N:         4,
		A:         -1,
		B:         3,
	}

	fmt.Println(algo.Solve("trapezoid", ig2))

	//server := web.New(cfg)
	//
	//_ = server.Start()
}
