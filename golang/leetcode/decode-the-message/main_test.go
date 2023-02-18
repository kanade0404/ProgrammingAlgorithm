package decode_the_message

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_decodeMessage(t *testing.T) {
	type args struct {
		key     string
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				key:     "the quick brown fox jumps over the lazy dog",
				message: "vkbs bs t suepuv",
			},
			want: "this is a secret",
		},
		{
			name: "test2",
			args: args{
				key:     "eljuxhpwnyrdgtqkviszcfmabo",
				message: "zwx hnfx lqantp mnoeius ycgk vcnjrdb",
			},
			want: "the five boxing wizards jump quickly",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decodeMessage(tt.args.key, tt.args.message)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
