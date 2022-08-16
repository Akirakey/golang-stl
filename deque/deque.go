package deque

import "sync"

const cache_SIZE = 1024

// Node of deque

type dequeNode[T any] struct {
	cache [cache_SIZE]T
	front int16
	back  int16
	next  *dequeNode[T]
	prev  *dequeNode[T]
}

func (node *dequeNode[T]) pushBack(value T) {
	node.cache[node.back] = value
	node.back += 1
}

func (node *dequeNode[T]) pushFront(value T) {
	node.cache[node.front] = value
	node.front -= 1
}

func (node *dequeNode[T]) popBack() (value T) {
	node.back -= 1
	value = node.cache[node.back]
	return value
}

func (node *dequeNode[T]) popFront() (value T) {
	node.front += 1
	value = node.cache[node.front]
	return value
}

// Deque

type Deque[T any] struct {
	first *dequeNode[T]
	last  *dequeNode[T]
	size  uint64
	mutex sync.Mutex
}

func (deque *Deque[T]) PushFront(value T) {
	deque.mutex.Lock()
	defer deque.mutex.Unlock()
	if deque.first.front < 0 {
		if deque.first.prev == nil {
			deque.first.prev = &dequeNode[T]{front: cache_SIZE - 1, back: cache_SIZE}
			deque.first.prev.next = deque.first
		}
		deque.first = deque.first.prev
	}
	deque.first.pushFront(value)
	deque.size += 1
}

func (deque *Deque[T]) PushBack(value T) {
	deque.mutex.Lock()
	defer deque.mutex.Unlock()
	if deque.last.back >= cache_SIZE {
		if deque.last.next == nil {
			deque.last.next = &dequeNode[T]{front: -1, back: 0}
			deque.last.next.prev = deque.last
		}
		deque.last = deque.last.next
	}
	deque.last.pushBack(value)
	deque.size += 1
}

func (deque *Deque[T]) PopFront() (value T) {
	if deque.Empty() {
		panic("Cannot pop an element from an empty Deque")
	}
	deque.mutex.Lock()
	defer deque.mutex.Unlock()
	if deque.first.front >= cache_SIZE-1 {
		deque.first = deque.first.next
		deque.first.prev = nil
	}
	deque.size -= 1
	return deque.first.popFront()
}

func (deque *Deque[T]) PopBack() (value T) {
	if deque.Empty() {
		panic("Cannot pop an element from an empty Deque")
	}
	deque.mutex.Lock()
	defer deque.mutex.Unlock()
	if deque.last.back <= 0 {
		deque.last = deque.last.prev
		deque.last.next = nil
	}
	deque.size -= 1
	return deque.last.popBack()
}

func (deque *Deque[T]) Size() uint64 {
	return deque.size
}

func (deque *Deque[T]) Empty() bool {
	return deque.size <= 0
}

func (deque *Deque[T]) Front() (value T) {
	if deque.Empty() {
		var value T
		return value
	}
	return deque.first.cache[deque.first.front+1]
}

func (deque *Deque[T]) Back() (value T) {
	if deque.Empty() {
		var value T
		return value
	}
	return deque.last.cache[deque.last.back-1]
}

func New[T any]() *Deque[T] {
	cache := &dequeNode[T]{front: cache_SIZE/2 - 1, back: cache_SIZE / 2}
	return &Deque[T]{
		first: cache,
		last:  cache,
	}
}
