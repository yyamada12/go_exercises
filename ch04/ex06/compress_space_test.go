package main

import (
	"reflect"
	"testing"
)

func Test_compressSpace(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"single fullwidth space", args{[]byte("hoge　fuga")}, []byte("hoge fuga")},
		{"fullwidth space and tab", args{[]byte("hoge\t　fuga")}, []byte("hoge fuga")},
		{"Begin with fullwidth space and tab", args{[]byte("　\thoge fuga")}, []byte(" hoge fuga")},
		{"End with fullwidth space and tab", args{[]byte("hoge\t　fuga\t ")}, []byte("hoge fuga ")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressSpace(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compressSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
