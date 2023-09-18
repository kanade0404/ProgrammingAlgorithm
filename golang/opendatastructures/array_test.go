package opendatastructures

import (
	"reflect"
	"testing"
)

func TestArrayStack_Get(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		a    ArrayStack[string]
		args args
		want string
	}{
		{
			name: "index0=a",
			a: ArrayStack[string]{
				arr: []string{"a", "b", "c"},
				n:   3,
			},
			args: args{
				index: 0,
			},
			want: "a",
		},
		{
			name: "index1=b",
			a: ArrayStack[string]{
				arr: []string{"a", "b", "c"},
				n:   3,
			},
			args: args{
				index: 1,
			},
			want: "b",
		},
		{
			name: "index2=c",
			a:    ArrayStack[string]{"a", "b", "c"},
			args: args{
				index: 2,
			},
			want: "c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Get(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		a    ArrayStack[string]
		want string
	}{
		{
			name: "test",
			a:    ArrayStack[string]{"a", "b", "c"},
			want: "c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayStack_Push(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		a    ArrayStack[string]
		args args
	}{
		{
			name: "",
			a:    ArrayStack[string]{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Push(tt.args.value)
		})
	}
}
