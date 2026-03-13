package numeric

import "math"

type Function struct {
	Name string
	Fn   func(float64) float64
}

var functions = []func(x float64) float64{
	func(x float64) float64 {
		return x*x*x + 2.84*x*x - 5.606*x - 14.766
	},
	func(x float64) float64 {
		return x*x*x - 1.89*x*x - 2*x + 1.76
	},
	func(x float64) float64 {
		return math.Sin(3*x) - 0.5
	},
}

func GetFunction(index int) func(float64) float64 {
	return functions[index]
}
