package day4

import (
	"reflect"
	"testing"
)

func TestMatrixReshape1(t *testing.T) {
	type args struct {
		mat [][]int
		r   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "TestMatrixReshape1-1", args: args{mat: [][]int{{1, 2}, {3, 4}}, r: 1, c: 4}, want: [][]int{{1, 2, 3, 4}}},
		{name: "TestMatrixReshape1-2", args: args{mat: [][]int{{1, 2}, {3, 4}}, r: 2, c: 4}, want: [][]int{{1, 2}, {3, 4}}},
		{name: "TestMatrixReshape1-3", args: args{mat: [][]int{{1, 2, 3, 4}}, r: 2, c: 2}, want: [][]int{{1, 2}, {3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatrixReshape1(tt.args.mat, tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatrixReshape1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrixReshape2(t *testing.T) {
	type args struct {
		mat [][]int
		r   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "TestMatrixReshape2-1", args: args{mat: [][]int{{1, 2}, {3, 4}}, r: 1, c: 4}, want: [][]int{{1, 2, 3, 4}}},
		{name: "TestMatrixReshape2-2", args: args{mat: [][]int{{1, 2}, {3, 4}}, r: 2, c: 4}, want: [][]int{{1, 2}, {3, 4}}},
		{name: "TestMatrixReshape2-3", args: args{mat: [][]int{{1, 2, 3, 4}}, r: 2, c: 2}, want: [][]int{{1, 2}, {3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatrixReshape2(tt.args.mat, tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatrixReshape2() = %v, want %v", got, tt.want)
			}
		})
	}
}
