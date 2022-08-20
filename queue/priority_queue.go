package queue

type PriorityQueue[T any] struct {
	arr  []T
	tail int
}

func NewPriorityQueue[T any](size int) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		arr:  make([]T, size),
		tail: 0,
	}
}
