package a_very_big_sum

import "testing"

func Test_aVeryBigSum(t *testing.T) {
	type args struct {
		ar []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"テスト",
			args{
				ar: []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005},
			},
			5000000015,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aVeryBigSum(tt.args.ar); got != tt.want {
				t.Errorf("aVeryBigSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
