package main

import (
	"comp-math-3/internal/algo"
	"comp-math-3/internal/numeric"
	"fmt"
)

func main() {
	ig := numeric.Integral{
		F:         numeric.GetFunction(0),
		Tolerance: 0.01,
		N:         10,
		A:         0,
		B:         2,
	}

	fmt.Println(algo.Solve("trapezoid", ig))
	//fmt.Println(algo.Solve("simpson", ig))

	//cfg, err := config.Get()
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//server := web.New(cfg)
	//
	//_ = server.Start()
}
