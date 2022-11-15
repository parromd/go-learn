package sorts

import (
	"reflect"
	"time"
)

func (s *Sort) InsertionSort() (int, time.Duration) {
	data := make([]int, len(s.Data))
	copy(data, s.Data)
	comparisons := 0
	swap := reflect.Swapper(data)

	start := time.Now()
	i := 1
	for i < len(data) {
		j := i
		for j > 0 && data[j-1] > data[j] {
			swap(j, j-1)
			j--
			comparisons++
		}
		i++
	}
	elapsed := time.Since(start)

	s.Comparisons = comparisons
	return comparisons, elapsed
}
