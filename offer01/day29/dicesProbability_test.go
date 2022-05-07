package day29

import (
	"reflect"
	"testing"
)

func Test_dicesProbability(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"Test_dicesProbability", struct{ n int }{n: 3}, []float64{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dicesProbability(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dicesProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
