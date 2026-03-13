package numeric

import "math"

type Function struct {
	Name string
	Fn   func(float64) float64
}

var functions = []func(x float64) float64{
	func(x float64) float64 {
		return -x*x*x - x*x + x + 3.0
	},
	func(x float64) float64 {
		return math.Exp(-x * x)
	},
	func(x float64) float64 {
		return math.Sin(10*x) * math.Exp(-x)
	},
}

func GetFunction(index int) func(float64) float64 {
	return functions[index]
}
