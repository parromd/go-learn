package sorts

import (
	"reflect"
	"time"
)

func (s *Sort) BubbleSort() (int, time.Duration) {
	data := make([]int, len(s.Data))
	copy(data, s.Data)
	comparisons := 0
	swap := reflect.Swapper(data)

	start := time.Now()
	swapped := true
	for swapped {
		swapped = false
		for idx, val := range data {
			if idx < len(data)-1 {
				if val > data[idx+1] {
					swap(idx, idx+1)
					swapped = true
				}
				comparisons++
			}
		}
	}
	elapsed := time.Since(start)

	s.Comparisons = comparisons
	return comparisons, elapsed
}
