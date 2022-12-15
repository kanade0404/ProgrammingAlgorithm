package goal_parser_interpretation

import "testing"

func Test_interpret(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				command: "G()(al)",
			},
			want: "Goal",
		},
		{
			name: "test2",
			args: args{
				command: "G()()()()(al)",
			},
			want: "Gooooal",
		},
		{
			name: "test3",
			args: args{
				command: "(al)G(al)()()G",
			},
			want: "alGalooG",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := interpret(tt.args.command); got != tt.want {
				t.Errorf("interpret() = %v, want %v", got, tt.want)
			}
		})
	}
}
