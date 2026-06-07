package parse

import (
	"testing"
)

func TestFind(t *testing.T) {

}

func TestParseFloat(t *testing.T) {

}

func TestParseSuffix(T *testing.T) {
	test := []struct {
		name  string
		input string
		wantF float64
		wantB bool
	}{
		{
			name:  "1024K",
			input: "1024K",
			wantF: 1024 * 1024,
			wantB: true,
		},
		{
			name:  "1M",
			input: "1M",
			wantF: 1024 * 1024,
			wantB: true,
		},
		{
			name:  "abc",
			input: "abc",
			wantF: 0,
			wantB: false,
		},
		{
			name:  "len(input) = 0",
			input: "",
			wantF: 0,
			wantB: false,
		},
	}
	for _, tt := range test {
		T.Run(tt.name, func(T *testing.T) {
			got, gotB := ParseSuffix(tt.input)
			if got != tt.wantF && gotB != tt.wantB {
				T.Errorf("ParseSuffix(%q) = %f, want (%f, %t)",
					tt.input,
					got,
					tt.wantF,
					tt.wantB)
			}
		})
	}
}
