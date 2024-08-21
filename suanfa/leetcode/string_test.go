package main

import (
	"reflect"
	"testing"
)

func TestFullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "t1", args: args{words: []string{"This", "is", "an", "example", "of", "text", "justification."}, maxWidth: 16}, want: []string{"This    is    an",
			"example  of text",
			"justification.  "}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
