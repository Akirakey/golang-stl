package queue

type PriorityQueue[T any] struct {
	arr     []T
	tail    int
	compare func(e1 T, e2 T) int
}

// The argument "compare" shoule be a function to compare two elements.
// When e1 > e2, this function should return a positive integer.
// When e1 == e2, this function should return 0.
// When e1 < e2, this function should return a negative integer.
// PriorityQueue uses ascending sort. if you want to reverse it, let the compare function return positive interger when e1 < e2.
func NewPriorityQueue[T any](size int, compare func(e1 T, e2 T) int) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		arr:     make([]T, size),
		tail:    0,
		compare: compare,
	}
}

func (pq *PriorityQueue[T]) add(item T) {
	if len(pq.arr) < cap(pq.arr) || pq.tail < len(pq.arr) {
		pq.arr[pq.tail] = item
	} else {
		pq.arr = append(pq.arr, item)
	}
	pq.tail++
}

func (pq *PriorityQueue[T]) Push(item T) {
	pq.add(item)
	if pq.tail <= 1 {
		return
	}
	if pq.tail == 2 {
		if pq.compare(pq.arr[0], pq.arr[1]) > 0 {
			pq.arr[0], pq.arr[1] = pq.arr[1], pq.arr[0]
		}
		return
	}
	var (
		child  int = pq.tail - 1
		parent int = (child - 1) / 2
	)
	for child != parent {
		if pq.compare(pq.arr[child], pq.arr[parent]) < 0 {
			pq.arr[parent], pq.arr[child] = pq.arr[child], pq.arr[parent]
		}
		child = parent
		parent = (child - 1) / 2
	}
}

func (pq *PriorityQueue[T]) sink() {
	// Has 1 of no element, do nothing
	if pq.tail <= 1 {
		return
	}
	// Has 2 elements, just compare them.
	if pq.tail == 2 {
		if pq.compare(pq.arr[0], pq.arr[1]) > 0 {
			pq.arr[0], pq.arr[1] = pq.arr[1], pq.arr[0]
		}
		return
	}
	var (
		parent int = 0
		child  int = 1
	)
	for child < pq.tail {
		// Choose the smaller one of two children
		if child+1 < pq.tail && pq.compare(pq.arr[child+1], pq.arr[child]) < 0 {
			child += 1
		}
		if pq.compare(pq.arr[child], pq.arr[parent]) < 0 {
			pq.arr[parent], pq.arr[child] = pq.arr[child], pq.arr[parent]
		}
		parent = child
		child = 2*parent + 1
	}

}

func (pq *PriorityQueue[T]) Pop() T {
	if pq.Size() == 0 {
		panic("Pop element from a empty queue.")
	}
	pq.tail--
	var res T = pq.arr[0]
	pq.arr[0] = pq.arr[pq.tail]
	pq.sink()
	return res
}

func (pq *PriorityQueue[T]) Head() T {
	return pq.arr[0]
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.tail
}
