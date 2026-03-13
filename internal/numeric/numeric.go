package numeric

type Solution struct {
	Value      float64
	Partitions int
}

type Integral struct {
	F         func(float64) float64
	Tolerance float64
	N         int
	A         float64
	B         float64
}
