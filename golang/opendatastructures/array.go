package opendatastructures

type ArrayStack[T any] struct {
	arr []T
	n   int
}

func New[T any]() *ArrayStack[T] {
	v := make([]T, 0)
	return &ArrayStack[T]{
		arr: v,
		n:   0,
	}
}
func (a *ArrayStack[T]) resize() {
	arrLen := len(a.arr)
	newArr := make([]T, arrLen*2)
	for i := 0; i < arrLen; i++ {
		newArr[i] = a.arr[i]
	}
	a.arr = newArr
}

func (a *ArrayStack[T]) Get(i int) T {
	return a.arr[i]
}

func (a *ArrayStack[T]) Set(i int, x T) T {
	y := a.arr[i]
	a.arr[i] = x
	a.n += 1
	return y
}

func (a *ArrayStack[T]) Add(i int, x T) {
	if a.n+1 > len(a.arr) {
		a.resize()
	}
	for j := a.n; j > i; j-- {
		a.arr[j] = a.arr[j-1]
	}
	a.arr[i] = x
	a.n += 1
}

func (a *ArrayStack[T]) Remove(i int) T {
	x := a.arr[i]
	for j := i; j < a.n; j++ {
		a.arr[j] = a.arr[j+1]
	}
	a.n -= 1
	if len(a.arr) >= 3*a.n {
		a.resize()
	}
	return x
}
