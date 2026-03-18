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
		return 1 / math.Sqrt(math.Abs(x))
	},
	func(x float64) float64 {
		return 1 / x
	},
	func(x float64) float64 {
		return 1 / math.Sqrt(2*x-x*x)
	},
}

func GetFunction(index int) func(float64) float64 {
	return functions[index]
}
