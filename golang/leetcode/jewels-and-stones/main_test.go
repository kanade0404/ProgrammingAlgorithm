package jewels_and_stones

import "testing"

func Test_numJewelsInStones(t *testing.T) {
	type args struct {
		jewels string
		stones string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				jewels: "aA",
				stones: "aAAbbbb",
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				jewels: "z",
				stones: "ZZ",
			},
			want: 0,
		},
		{
			name: "test3",
			args: args{
				jewels: "bB",
				stones: "bBBb",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStones(tt.args.jewels, tt.args.stones); got != tt.want {
				t.Errorf("numJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
