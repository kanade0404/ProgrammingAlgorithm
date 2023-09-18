package algorithm_introduction

type InsertSort struct {
	arr []int
}

func (is *InsertSort) Sort() {
	l := len(is.arr)
	for j := 1; j < l; j++ {
		key := is.arr[j]
		i := j - 1
		for i > 0 && is.arr[i] > key {
			is.arr[i+1] = is.arr[i]
			i = i - 1
		}
		is.arr[i+1] = key
	}
}
