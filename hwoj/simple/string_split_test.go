package simple

import "testing"

func TestStringSplit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "TestStringSplit-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StringSplit()
		})
	}
}
