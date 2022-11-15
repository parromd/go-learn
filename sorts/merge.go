package sorts

import (
	"container/list"
	"time"
)

// Queue is a queue
type Queue interface {
	Front() *list.Element
	Len() int
	Add(interface{})
	Remove() interface{}
}

type queueImpl struct {
	*list.List
}

func (q queueImpl) Len() int {
	return q.List.Len()
}

func (q *queueImpl) Add(v interface{}) {
	q.PushBack(v)
}

func (q *queueImpl) Remove() interface{} {
	if q.Len() == 0 {
		return []int{}
	}
	e := q.Front()
	q.List.Remove(e)
	return e.Value
}

// New is a new instance of a Queue
func NewQ() Queue {
	return &queueImpl{list.New()}
}

func (s *Sort) MergeSort() (int, time.Duration) {
	data := make([]int, len(s.Data))
	copy(data, s.Data)
	comparisons := 0

	start := time.Now()
	q := NewQ()
	for _, val := range data {
		q.Add([]int{val})
	}
	var merge func([]int, []int) []int
	merge = func(l, r []int) []int {
		if len(l) == 0 {
			return r
		}

		if len(r) == 0 {
			return l
		}

		comparisons++
		if l[0] <= r[0] {
			return append(l[:1], merge(l[1:], r)...)
		}
		return append(r[:1], merge(l, r[1:])...)
	}

	for q.Len() > 1 {
		q.Add(merge(q.Remove().([]int), q.Remove().([]int)))
	}

	elapsed := time.Since(start)
	s.Comparisons = comparisons
	return comparisons, elapsed
}
