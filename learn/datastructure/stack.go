package datastructure

type GoStack[T any] struct {
	container []T
	size      int32
}

func New[T any]() *GoStack[T] {
	return &GoStack[T]{
		container: []T{},
		size:      0,
	}
}

func (goStack *GoStack[T]) Push(v T) {
	goStack.container = append(goStack.container, v)
	goStack.size++
}

func (goStack *GoStack[T]) Pop() (v T) {
	if goStack.size == 0 {
		panic(v)
		return
	}
	v = goStack.container[goStack.size-1]
	goStack.container = goStack.container[:goStack.size-1]
	goStack.size--
	return
}

func (goStack *GoStack[T]) Peek() (v T) {
	if goStack.size == 0 {
		panic(v)
		return v
	}

	v = goStack.container[goStack.size-1]
	return
}

func (goStack *GoStack[T]) Size() int32 {
	return goStack.size
}
