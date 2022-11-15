package sorts

// I wanted to make this generic over any `comparable` type but
// doing so isn't quite so easy in go atm (plus I'm dumb)

type Sort struct {
	capacity    int
	Data        []int
	Comparisons int
}

func (s *Sort) Set(new []int) {
	copy(s.Data, new)
}

func New(capacity int) *Sort {
	return &Sort{capacity: capacity, Data: make([]int, capacity)}
}
