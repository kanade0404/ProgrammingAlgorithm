package design_an_ordered_stream

type OrderedStream struct {
	Arr      []string
	Len      int
	StartIdx int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		Arr: make([]string, n, n),
		Len: n,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.Arr[idKey-1] = value
	results := make([]string, 0)
	count := this.StartIdx
	for i := count; i < this.Len; i++ {
		if this.Arr[i] == "" {
			break
		}
		results = append(results, this.Arr[i])
		this.StartIdx = i + 1
	}
	return results
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
