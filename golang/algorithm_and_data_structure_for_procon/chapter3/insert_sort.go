package chapter3

type Arr struct {
	Nums []int
}

func (a *Arr) InsertSort(n int) {
	for i := 1; i < n; i++ {
		v := a.Nums[i]
		j := i - 1
		for j >= 0 && a.Nums[j] > v {
			a.Nums[j+1] = a.Nums[j]
			j = j - 1
		}
		a.Nums[j+1] = v
	}
}
