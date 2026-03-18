package algo

import (
	"comp-math-3/internal/numeric"
	"errors"
	"fmt"
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

const epsOffset = 1e-6

func solveConvergent(method string, ig numeric.Integral) (numeric.Solution, error) {
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

func Solve(method string, ig numeric.Integral) (numeric.Solution, error) {
	a, b := ig.A, ig.B
	f := ig.F

	breakpoints := GetDiscontinuityPoints(ig, int(math.Ceil(b-a))*1_000)

	fmt.Println("breakpoints:", breakpoints)

	if len(breakpoints) != 0 {
		// Проверка сходимости
		converges := true

		for _, bp := range breakpoints {
			y1 := TryToCompute(f, bp-epsOffset)
			y2 := TryToCompute(f, bp+epsOffset)

			// Проверяем разные случаи расходимости
			if y1 == nil && y2 == nil {
				// Точка неустранимого разрыва (например, 1/x в 0)
				// Проверяем поведение рядом с точкой
				y1Near := TryToCompute(f, bp-epsOffset)
				y2Near := TryToCompute(f, bp+epsOffset)
				if y1Near == nil || y2Near == nil {
					converges = false
					break
				}
			} else {
				// С одной стороны функция существует, с другой - нет
				// Это может быть устранимый разрыв или особенность типа 1/sqrt(x)
				// Дополнительная проверка
				var yExist *float64
				if y1 != nil {
					yExist = y1
				} else {
					yExist = y2
				}

				// Проверяем, стремится ли функция к бесконечности
				if math.IsInf(*yExist, 0) {
					converges = false
					break
				}
			}
		}

		if !converges {
			return numeric.Solution{}, errors.New("integral is not convergent on the interval " +
				"(has non-integrable singularity)")
		}

		// Если интеграл сходится, обрабатываем особые точки
		if len(breakpoints) == 1 {
			fmt.Println("breakpoints:", len(breakpoints))
			singularity := breakpoints[0]

			// Особенность на левой границе
			if math.Abs(singularity-a) < epsOffset {
				ig.A = a + epsOffset
				return solveConvergent(method, ig)
			} else if math.Abs(singularity-b) < epsOffset {
				// Особенность на правой границе
				ig.B = b - epsOffset
				return solveConvergent(method, ig)
			} else {
				// Особенность внутри интервала: разбиваем на две части
				ig1 := ig
				ig1.B = singularity - epsOffset
				res1, err1 := solveConvergent(method, ig1)
				if err1 != nil {
					return numeric.Solution{}, fmt.Errorf("error on left part: %w", err1)
				}

				ig2 := ig
				ig2.A = singularity + epsOffset
				res2, err2 := solveConvergent(method, ig2)
				if err2 != nil {
					return numeric.Solution{}, fmt.Errorf("error on right part: %w", err2)
				}

				return numeric.Solution{
					Value:      res1.Value + res2.Value,
					Partitions: res1.Partitions + res2.Partitions,
				}, nil
			}
		} else {
			// Несколько точек разрыва
			// Сортируем точки разрыва
			breakpointsSorted := make([]float64, len(breakpoints))
			copy(breakpointsSorted, breakpoints)
			for i := 0; i < len(breakpointsSorted)-1; i++ {
				for j := i + 1; j < len(breakpointsSorted); j++ {
					if breakpointsSorted[i] > breakpointsSorted[j] {
						breakpointsSorted[i], breakpointsSorted[j] = breakpointsSorted[j], breakpointsSorted[i]
					}
				}
			}

			totalValue := 0.0
			totalPartitions := 0

			// Интервал от a до первой точки разрыва
			if a < breakpointsSorted[0]-epsOffset {
				igSeg := ig
				igSeg.B = breakpointsSorted[0] - epsOffset
				res, err := solveConvergent(method, igSeg)
				if err != nil {
					return numeric.Solution{}, fmt.Errorf("error on segment [%f, %f]: %w", a, breakpointsSorted[0]-epsOffset, err)
				}
				totalValue += res.Value
				totalPartitions += res.Partitions
			}

			// Интервалы между точками разрыва
			for i := 0; i < len(breakpointsSorted)-1; i++ {
				if breakpointsSorted[i]+epsOffset < breakpointsSorted[i+1]-epsOffset {
					igSeg := ig
					igSeg.A = breakpointsSorted[i] + epsOffset
					igSeg.B = breakpointsSorted[i+1] - epsOffset
					res, err := solveConvergent(method, igSeg)
					if err != nil {
						return numeric.Solution{}, fmt.Errorf("error on segment [%f, %f]: %w",
							breakpointsSorted[i]+epsOffset, breakpointsSorted[i+1]-epsOffset, err)
					}
					totalValue += res.Value
					totalPartitions += res.Partitions
				}
			}

			// Интервал от последней точки разрыва до b
			if breakpointsSorted[len(breakpointsSorted)-1]+epsOffset < b {
				igSeg := ig
				igSeg.A = breakpointsSorted[len(breakpointsSorted)-1] + epsOffset
				res, err := solveConvergent(method, igSeg)
				if err != nil {
					return numeric.Solution{}, fmt.Errorf("error on segment [%f, %f]: %w",
						breakpointsSorted[len(breakpointsSorted)-1]+epsOffset, b, err)
				}
				totalValue += res.Value
				totalPartitions += res.Partitions
			}

			return numeric.Solution{
				Value:      totalValue,
				Partitions: totalPartitions,
			}, nil
		}
	}

	return solveConvergent(method, ig)
}

func calcRunge(method string, i0 float64, i1 float64) float64 {
	return math.Abs(i1-i0) / (math.Pow(2, orderOfAccuracy[method]) - 1)
}
