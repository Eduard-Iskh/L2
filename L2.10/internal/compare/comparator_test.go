package compare

import (
	"l210-sort/internal/config"

	"testing"
)

func TestCompareSring(t *testing.T) {

	test := []struct {
		name   string
		inputA config.LineComp
		inputB config.LineComp
		cfg    config.Config
		want   int
		ok     bool
	}{{
		name: "Test 1: a < b",
		inputA: config.LineComp{
			CompElemS: "apple",
		},
		inputB: config.LineComp{
			CompElemS: "banana",
		},
		cfg:  config.Config{},
		want: -1,
		ok:   true,
	},
		{
			name: "Test 2: a > b",
			inputA: config.LineComp{
				CompElemS: "banana",
			},
			inputB: config.LineComp{
				CompElemS: "apple",
			},
			cfg:  config.Config{},
			want: 1,
			ok:   true,
		},
		{
			name: "Test 2: a = b",
			inputA: config.LineComp{
				CompElemS: "apple",
			},
			inputB: config.LineComp{
				CompElemS: "apple",
			},
			cfg:  config.Config{},
			want: 0,
			ok:   true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := Compare(tt.inputA, tt.inputB, tt.cfg)
			if got != tt.want {
				t.Fatalf(
					"Compare(%q, %q) = %d, want %d",
					tt.inputA.CompElemS,
					tt.inputB.CompElemS,
					got,
					tt.want)
			}
		})
	}
}

func TestCompareNumbers(T *testing.T) {
	test := []struct {
		name   string
		inputA config.LineComp
		inputB config.LineComp
		want   int
	}{{
		name: "Test 1: a < b",
		inputA: config.LineComp{
			FCompElemS: 1,
			IsNumber:   true,
		},
		inputB: config.LineComp{
			FCompElemS: 2,
			IsNumber:   true,
		},
		want: -1,
	},
		{
			name: "Test 2: a > b",
			inputA: config.LineComp{
				FCompElemS: 10,
				IsNumber:   true,
			},
			inputB: config.LineComp{
				FCompElemS: 2,
				IsNumber:   true,
			},
			want: 1,
		},
		{
			name: "Test 3: a = b",
			inputA: config.LineComp{
				FCompElemS: 0,
				IsNumber:   true,
			},
			inputB: config.LineComp{
				FCompElemS: 0,
				IsNumber:   true,
			},
			want: 0,
		},
	}
	cfg := config.Config{
		N: true,
	}
	for _, tt := range test {
		T.Run(tt.name, func(T *testing.T) {
			got := Compare(tt.inputA, tt.inputB, cfg)
			if got != tt.want {
				T.Fatalf("Compare(%f, %f) = %d, want %d",
					tt.inputA.FCompElemS,
					tt.inputB.FCompElemS,
					got,
					tt.want)
			}
		})
	}
}

func TestCompareSuffix(T *testing.T) {
	test := []struct {
		name   string
		inputA config.LineComp
		inputB config.LineComp
		want   int
	}{
		{
			name:   "Test 1: 1K < 2K",
			inputA: config.LineComp{CompElemS: "1K"},
			inputB: config.LineComp{CompElemS: "2K"},
			want:   -1},
		{
			name:   "Test 2: 1M > 500K",
			inputA: config.LineComp{CompElemS: "1M"},
			inputB: config.LineComp{CompElemS: "500K"},
			want:   1,
		},
		{
			name:   "Test 3: 1024K == 1M",
			inputA: config.LineComp{CompElemS: "1024K"},
			inputB: config.LineComp{CompElemS: "1M"},
			want:   0,
		},
		{
			name:   "Test 4: 1K vs abc",
			inputA: config.LineComp{CompElemS: "1K"},
			inputB: config.LineComp{CompElemS: "abc"},
			want:   1,
		},
		{
			name:   "Test 5: abc vs xyz",
			inputA: config.LineComp{CompElemS: "abc"},
			inputB: config.LineComp{CompElemS: "xyz"},
			want:   -1,
		},
	}
	cfg := config.Config{
		H: true,
	}
	for _, tt := range test {
		T.Run(tt.name, func(T *testing.T) {
			got := Compare(tt.inputA, tt.inputB, cfg)
			if got != tt.want {
				T.Errorf("Compare(%q, %q) = %d, want %d",
					tt.inputA.CompElemS,
					tt.inputB.CompElemS,
					got,
					tt.want)
			}
		})
	}
}
func TestLess(t *testing.T) {

}
