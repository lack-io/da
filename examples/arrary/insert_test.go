package arrary

import (
	"reflect"
	"testing"
)

func TestInsert1(t *testing.T) {
	type args struct {
		list  []int
		index int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test-1",
			args{
				list:  []int{1, 3, 4},
				index: 1,
				n:     2,
			},
			[]int{1, 2, 3, 4},
		},
		{"Test-2",
			args{
				list:  []int{1, 2, 3},
				index: 3,
				n:     4,
			},
			[]int{1, 2, 3, 4},
		},
		{"Test-3",
			args{
				list:  []int{1, 2, 3},
				index: 0,
				n:     0,
			},
			[]int{0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.args.list, tt.args.index, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		list  []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"TestRemove-1",
			args{
				list:  []int{1, 2, 3},
				index: 0,
			},
			[]int{2, 3},
		},
		{
			"TestRemove-2",
			args{
				list:  []int{1, 2, 3},
				index: 1,
			},
			[]int{1, 3},
		},
		{
			"TestRemove-1",
			args{
				list:  []int{1, 2, 3},
				index: 2,
			},
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.list, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
