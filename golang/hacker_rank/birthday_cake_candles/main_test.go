package birthday_cake_candles

import "testing"

func Test_birthdayCakeCandles(t *testing.T) {
	type args struct {
		candles []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			"テスト",
			args{
				[]int32{3, 2, 1, 3},
			},
			2,
		},
		{
			"全て一種類",
			args{
				[]int32{5, 3, 2, 1, 4},
			},
			1,
		},
		{
			"全て同じ",
			args{
				[]int32{1, 1, 1, 1},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := birthdayCakeCandles(tt.args.candles); got != tt.want {
				t.Errorf("birthdayCakeCandles() = %v, want %v", got, tt.want)
			}
		})
	}
}
