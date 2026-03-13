package algo

import (
	"comp-math-3/internal/numeric"
	"errors"
	"math"
)

var methods = map[string]func(integral numeric.Integral) (float64, error){
	"rectangleLeft":     SolveLeftRectangle,
	"rectangleRight":    SolveRightRectangle,
	"rectangleMidpoint": SolveMidpointRectangle,
	"trapezoid":         SolveTrapezoid,
	"simpson":           SolveSimpson,
}

var orderOfAccuracy = map[string]float64{
	"rectangleLeft":     1,
	"rectangleRight":    1,
	"rectangleMidpoint": 2,
	"trapezoid":         2,
	"simpson":           4,
}

func Solve(method string, ig numeric.Integral) (numeric.Solution, error) {
	if ig.A >= ig.B {
		return numeric.Solution{}, errors.New("a must be higher than b")
	}

	if ig.Tolerance <= 0 {
		return numeric.Solution{}, errors.New("eps must be greater than zero")
	}

	if _, ok := methods[method]; !ok {
		return numeric.Solution{}, errors.New("unknown method")
	}

	val0, err0 := methods[method](ig)
	if err0 != nil {
		return numeric.Solution{}, err0
	}

	ig.N = ig.N * 2

	val1, err1 := methods[method](ig)
	if err1 != nil {
		return numeric.Solution{}, err1
	}

	r := calcRunge(method, val0, val1)

	for r > ig.Tolerance {
		ig.N = ig.N * 2
		temp := val1
		val1, err1 = methods[method](ig)
		if err1 != nil {
			return numeric.Solution{}, err1
		}
		val0 = temp

		r = calcRunge(method, val0, val1)
	}

	return numeric.Solution{
		Value:      val1,
		Partitions: ig.N,
	}, nil
}

func calcRunge(method string, i0 float64, i1 float64) float64 {
	return math.Abs(i1-i0) / (math.Pow(2, orderOfAccuracy[method]) - 1)
}
