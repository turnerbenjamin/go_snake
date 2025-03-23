package utilities

type List[T any] interface {
	Push(item T)
	Size() int
	Get(int) T
	Clear()
}

type simpleList[T any] struct {
	items []T
	size  int
}

func NewList[T any](maxSize int) List[T] {
	return &simpleList[T]{
		items: make([]T, maxSize),
		size:  0,
	}
}

func (l *simpleList[T]) Push(item T) {
	l.items[l.size] = item
	l.size++
}

func (l *simpleList[T]) Size() int {
	return l.size
}

func (l *simpleList[T]) Get(i int) T {
	return l.items[i]
}

func (l *simpleList[T]) Clear() {
	l.size = 0
}