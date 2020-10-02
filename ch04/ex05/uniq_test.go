package main

import (
	"reflect"
	"testing"
)

func Test_uniq(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"only one", args{[]string{"hoge"}}, []string{"hoge"}},
		{"same words", args{[]string{"hoge", "hoge", "hoge"}}, []string{"hoge"}},
		{"many words", args{[]string{"hoge", "piyo", "piyo"}}, []string{"hoge", "piyo"}},
		{"empty", args{[]string{}}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniq(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
