package day9

import "testing"

func TestIsValid1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "TestIsValid1-1", args: struct{ s string }{s: "([])"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid1(tt.args.s); got != tt.want {
				t.Errorf("IsValid1() = %v, want %v", got, tt.want)
			}
		})
	}
}
