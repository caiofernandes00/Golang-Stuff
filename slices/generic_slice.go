package slices

type GenericSlice[T comparable] struct {
	data []T
}

func New[T comparable]() *GenericSlice[T] {
	return &GenericSlice[T]{
		data: []T{},
	}
}

func (l *GenericSlice[T]) Insert(value T) {
	l.data = append(l.data, value)
}

func (l *GenericSlice[T]) Get(i int) T {
	if i > len(l.data)-1 {
		panic("given index is too high")
	}

	for it := 0; it < len(l.data); it++ {
		if i == it {
			return l.data[it]
		}
	}

	panic("value not found")
}

func (l *GenericSlice[T]) RemoveByIndex(i int) {
	if i > len(l.data)-1 {
		panic("given index is too high")
	}

	for it := 0; it < len(l.data); it++ {
		if it == i {
			l.data = append(l.data[:it], l.data[it+1:]...)
		}
	}
}

func (l *GenericSlice[T]) RemoveByValue(value T) {
	for i := 0; i < len(l.data); i++ {
		if value == l.data[i] {
			l.data = append(l.data[:i], l.data[i+1:]...)
		}
	}
}
