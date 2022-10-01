package sorts

import (
	"time"
	"reflect"
	"fmt"
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
			if idx < len(data) - 1 {
				if val > data[idx + 1] {
					swap(idx, idx + 1)
					swapped = true
				}
				comparisons++
			}
		}
	}
	elapsed := time.Now().Sub(start)

	fmt.Println(data)
	s.Comparisons = comparisons
	return comparisons, elapsed
}