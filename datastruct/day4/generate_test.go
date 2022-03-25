package day4

import (
	"reflect"
	"testing"
)

func TestGenerate1(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "TestGenerate1-1", args: struct{ numRows int }{numRows: 5}, want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{name: "TestGenerate2-1", args: struct{ numRows int }{numRows: 1}, want: [][]int{{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Generate1(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate1() = %v, want %v", got, tt.want)
			}
		})
	}
}
