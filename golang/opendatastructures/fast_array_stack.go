package opendatastructures

type FastArrayStack[T any] []T

func (f *FastArrayStack[T]) resize() {
	newArr := make(FastArrayStack[T], len(*f)*2)
	copy(newArr, *f)
	*f = newArr
}
