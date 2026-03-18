package algo

import (
	"comp-math-3/internal/numeric"
	"math"
	"sort"
)

const (
	absoluteJumpThreshold = 1e8    // Абсолютный скачок
	relativeJumpThreshold = 1000.0 // Относительный скачок (в раз)
)

// TryToCompute безопасно вычисляет функцию, обрабатывая панику и особые значения.
func TryToCompute(f func(float64) float64, x float64) *float64 {
	defer func() {
		if r := recover(); r != nil {
			// паника при вычислении (например, деление на ноль) -> разрыв
		}
	}()

	result := f(x)
	if math.IsNaN(result) || math.IsInf(result, 0) {
		return nil
	}
	return &result
}

// GetDiscontinuityPoints возвращает отсортированный список точек,
// в которых функция f имеет разрыв (или подозрение на разрыв) на интервале [a,b].
// Параметр n задаёт количество равномерных шагов для первичного анализа.
func GetDiscontinuityPoints(ig numeric.Integral, n int) []float64 {
	f := ig.F
	a := ig.A
	b := ig.B

	h := (b - a) / float64(n)

	// Карта для сбора точек-кандидатов (ключ – средняя точка интервала, где заподозрен разрыв)
	candidates := make(map[float64]bool)

	// Значение в предыдущей точке
	prevX := a
	prevVal := TryToCompute(f, a)
	if prevVal == nil {
		// Если функция не определена на левой границе, сразу фиксируем разрыв в a
		candidates[a] = true
	}

	for i := 1; i <= n; i++ {
		x := a + float64(i)*h
		currVal := TryToCompute(f, x)

		// Если текущая точка не определена – разрыв в самой точке
		if currVal == nil {
			candidates[x] = true
			prevX = x
			prevVal = nil
			continue
		}

		// Если предыдущая точка была не определена – разрыв уже зафиксирован, просто обновляем
		if prevVal == nil {
			prevX = x
			prevVal = currVal
			continue
		}

		// Оба значения определены – проверяем, нет ли между ними резкого скачка
		diff := math.Abs(*currVal - *prevVal)
		maxAbs := math.Max(math.Abs(*prevVal), math.Abs(*currVal))

		// Избегаем деления на ноль при малых значениях
		if maxAbs < 1e-12 {
			maxAbs = 1e-12
		}

		if diff > absoluteJumpThreshold || diff > relativeJumpThreshold*maxAbs {
			// Разрыв где-то между prevX и x – добавляем середину интервала
			mid := (prevX + x) / 2
			candidates[mid] = true
		}

		prevX = x
		prevVal = currVal
	}

	// Преобразуем карту в срез и сортируем
	points := make([]float64, 0, len(candidates))
	for p := range candidates {
		points = append(points, p)
	}
	sort.Float64s(points)

	// Кластеризация: объединяем точки, находящиеся ближе, чем шаг h
	if len(points) == 0 {
		return points
	}

	eps := h // точность группировки (можно регулировать)
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
		// Берем среднее арифметическое кластера
		clustered = append(clustered, sum/float64(count))
		start = end + 1
	}

	return clustered
}
