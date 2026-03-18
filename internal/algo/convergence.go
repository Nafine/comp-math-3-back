package algo

import (
	"comp-math-3/internal/numeric"
	"math"
	"sort"
)

const (
	absoluteJumpThreshold = 1e7
	relativeJumpThreshold = 1000.0
)

func TryToCompute(f func(float64) float64, x float64) *float64 {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	result := f(x)
	if math.IsNaN(result) || math.IsInf(result, 0) {
		return nil
	}
	return &result
}

func GetDiscontinuityPoints(ig numeric.Integral, n int) []float64 {
	f := ig.F
	a := ig.A
	b := ig.B

	h := (b - a) / float64(n)

	candidates := make(map[float64]bool)

	prevX := a
	prevVal := TryToCompute(f, a)
	if prevVal == nil {
		candidates[a] = true
	}

	for i := 1; i <= n; i++ {
		x := a + float64(i)*h
		currVal := TryToCompute(f, x)

		if currVal == nil {
			candidates[x] = true
			prevX = x
			prevVal = nil
			continue
		}

		if prevVal == nil {
			prevX = x
			prevVal = currVal
			continue
		}

		diff := math.Abs(*currVal - *prevVal)
		maxAbs := math.Max(math.Abs(*prevVal), math.Abs(*currVal))

		if maxAbs < 1e-12 {
			maxAbs = 1e-12
		}

		if diff > absoluteJumpThreshold || diff > relativeJumpThreshold*maxAbs {
			mid := (prevX + x) / 2
			candidates[mid] = true
		}

		prevX = x
		prevVal = currVal
	}

	points := make([]float64, 0, len(candidates))
	for p := range candidates {
		points = append(points, p)
	}
	sort.Float64s(points)

	if len(points) == 0 {
		return points
	}

	eps := h
	clustered := make([]float64, 0, len(points))

	start := 0
	for start < len(points) {
		end := start
		sum := points[start]
		count := 1
		for end+1 < len(points) && points[end+1]-points[start] <= eps {
			end++
			sum += points[end]
			count++
		}
		clustered = append(clustered, sum/float64(count))
		start = end + 1
	}

	return clustered
}
