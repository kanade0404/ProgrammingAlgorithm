package first_letter_to_appear_twice

import "testing"

func Test_repeatedCharacter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "test1",
			args: args{
				s: "abccbaacz",
			},
			want: 'c',
		},
		{
			name: "test2",
			args: args{
				s: "abcdd",
			},
			want: 'd',
		},
		{
			name: "test3",
			args: args{
				s: "ncwn",
			},
			want: 'n',
		},
		{
			name: "test4",
			args: args{
				s: "aa",
			},
			want: 'a',
		},
		{
			name: "test5",
			args: args{
				s: "eessl",
			},
			want: 'e',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedCharacter(tt.args.s); got != tt.want {
				t.Errorf("repeatedCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
