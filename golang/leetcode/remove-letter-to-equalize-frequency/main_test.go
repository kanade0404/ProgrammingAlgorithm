package remove_letter_to_equalize_frequency

import "testing"

func Test_equalFrequency(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				word: "abcc",
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				word: "aazz",
			},
			want: false,
		},
		{
			name: "singles",
			args: args{
				word: "abc",
			},
			want: true,
		},
		{
			name: "2+3",
			args: args{
				word: "aabbb",
			},
			want: true,
		},
		{
			name: "2+2+1+1",
			args: args{
				"ddaccb",
			},
			want: false,
		},
		{
			name: "1+3",
			args: args{
				"abbb",
			},
			want: true,
		},
		{
			name: "1+2+2",
			args: args{
				"abbcc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalFrequency(tt.args.word); got != tt.want {
				t.Errorf("equalFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
