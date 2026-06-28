package grep

import (
	"slices"
	"testing"

	"example.com/internal/cfg"
	"example.com/internal/match"
)

func TestRun(t *testing.T) {

	test := []struct {
		name      string
		config    cfg.Config
		data      []string
		wantMerge [][2]int
		wantCount int
		wantErr   bool
	}{
		{
			name: "basic regexp",
			config: cfg.Config{
				Pattern: "apple",
			},
			data: []string{
				"apple",
				"banana",
				"adfg apple",
			},
			wantMerge: [][2]int{{0, 0}, {2, 2}},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "fixed string",
			config: cfg.Config{
				Pattern: "a.c",
				Fixed:   true,
			},
			data: []string{
				"a.c",
				"abc",
				"x a.c y",
			},
			wantMerge: [][2]int{{0, 0}, {2, 2}},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "ignore case",
			config: cfg.Config{
				Pattern:    "apple",
				IgnoreCase: true,
			},
			data: []string{
				"Apple",
				"banana",
				"APPLE",
			},
			wantMerge: [][2]int{{0, 0}, {2, 2}},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "invert filter",
			config: cfg.Config{
				Pattern: "apple",
				Invert:  true,
			},
			data: []string{
				"apple",
				"banana",
				"pear",
			},
			wantMerge: [][2]int{{1, 2}},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "context C",
			config: cfg.Config{
				Pattern: "banana",
				Context: 1,
			},
			data: []string{
				"apple",
				"banana",
				"pear",
			},
			wantMerge: [][2]int{{0, 2}},
			wantCount: 1,
			wantErr:   false,
		},
		{
			name: "after A",
			config: cfg.Config{
				Pattern: "banana",
				After:   1,
			},
			data: []string{
				"apple",
				"banana",
				"pear",
			},
			wantMerge: [][2]int{{1, 2}},
			wantCount: 1,
			wantErr:   false,
		},
		{
			name: "before B",
			config: cfg.Config{
				Pattern: "banana",
				Before:  1,
			},
			data: []string{
				"apple",
				"banana",
				"pear",
			},
			wantMerge: [][2]int{{0, 1}},
			wantCount: 1,
			wantErr:   false,
		},
		{
			name: "merge overlapping contexts",
			config: cfg.Config{
				Pattern: "banana|pear",
				Context: 1,
			},
			data: []string{
				"apple",
				"banana",
				"pear",
				"orange",
			},
			wantMerge: [][2]int{{0, 3}},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "invalid regexp",
			config: cfg.Config{
				Pattern: "(",
			},
			data: []string{
				"apple",
			},
			wantMerge: nil,
			wantCount: 0,
			wantErr:   true,
		},
		{
			name: "empty input",
			config: cfg.Config{
				Pattern: "apple",
			},
			data:      []string{},
			wantMerge: nil,
			wantCount: 0,
			wantErr:   false,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			merge, countMatches, err := Run(tt.data, tt.config, match.IsMatch)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !slices.Equal(merge, tt.wantMerge) {
				t.Fatalf("Run() merge = %v, want %v", merge, tt.wantMerge)
			}
			if countMatches != tt.wantCount {
				t.Fatalf("Run() count = %d, want %d", countMatches, tt.wantCount)
			}
		})
	}
}
