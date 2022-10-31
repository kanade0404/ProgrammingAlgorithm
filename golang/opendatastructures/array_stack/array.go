package array_stack

type ArrayStack[T any] []T

func New[T any]() *ArrayStack[T] {
	v := make(ArrayStack[T], 0)
	return &v
}

func (a *ArrayStack[T]) Get(index int) T {
	return (*a)[index]
}
func (a *ArrayStack[T]) resize() {
	newArr := make(ArrayStack[T], len(*a)*2)
	arrLen := len(*a)
	for i := 0; i < arrLen; i++ {
		newArr[i] = (*a)[i]
	}
	a = &newArr
}
func (a *ArrayStack[T]) Push(value T) {
	*a = append(*a, value)
}
func (a *ArrayStack[T]) Pop() T {
	v := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return v
}
