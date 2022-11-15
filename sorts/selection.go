package sorts

import (
	"reflect"
	"time"
)

func (s *Sort) SelectionSort() (int, time.Duration) {
	data := make([]int, len(s.Data))
	copy(data, s.Data)
	comparisons := 0
	swap := reflect.Swapper(data)

	start := time.Now()
	for i := 0; i < len(data)-1; i++ {
		min_idx := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[min_idx] {
				min_idx = j
			}
			comparisons++
		}
		if min_idx != i {
			swap(i, min_idx)
		}
	}
	elapsed := time.Since(start)

	s.Comparisons = comparisons
	return comparisons, elapsed
}
