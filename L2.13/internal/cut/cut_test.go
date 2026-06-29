package cut

import (
	"reflect"
	"slices"
	"testing"

	"example.com/internal/cfg"
)

func TestParseField(t *testing.T) {
	test := []struct {
		name     string
		line     []string
		fields   [][2]int
		wantData []string
	}{
		{
			name:     "single field",
			line:     []string{"a", "b", "c"},
			fields:   [][2]int{{1, 1}},
			wantData: []string{"a"},
		},
		{
			name:     "single range",
			line:     []string{"a", "b", "c", "d"},
			fields:   [][2]int{{2, 4}},
			wantData: []string{"b", "c", "d"},
		},
		{
			name:     "mixed single and range",
			line:     []string{"a", "b", "c", "d", "e"},
			fields:   [][2]int{{1, 1}, {3, 4}},
			wantData: []string{"a", "c", "d"},
		},
		{
			name:     "out of bounds single field is ignored",
			line:     []string{"a", "b"},
			fields:   [][2]int{{3, 3}},
			wantData: []string{},
		},
		{
			name:     "out of bounds after valid field should not stop later valid ranges",
			line:     []string{"a", "b"},
			fields:   [][2]int{{1, 1}, {5, 5}, {2, 2}},
			wantData: []string{"a", "b"},
		},
		{
			name:     "range partially out of bounds",
			line:     []string{"a", "b", "c"},
			fields:   [][2]int{{2, 5}},
			wantData: []string{"b", "c"},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			data := parseFields(tt.line, tt.fields)
			if !slices.Equal(data, tt.wantData) {
				t.Fatalf("parseField() = %v, want %v", data, tt.wantData)
			}
		})
	}
}

func TestCut(t *testing.T) {
	test := []struct {
		name   string
		data   []string
		config cfg.Config
		want   [][]string
	}{
		{
			name: "default delimiter",
			data: []string{"a\tb\tc"},
			config: cfg.Config{
				Fields:    [][2]int{{1, 1}, {3, 3}},
				Delimiter: "\t",
				Separated: false,
			},
			want: [][]string{{"a", "c"}},
		},
		{
			name: "custom delimiter",
			data: []string{"a,b,c"},
			config: cfg.Config{
				Fields:    [][2]int{{2, 2}},
				Delimiter: ",",
				Separated: false,
			},
			want: [][]string{{"b"}},
		},
		{
			name: "separated skips line without delimiter",
			data: []string{"a,b,c", "hello", "x,y"},
			config: cfg.Config{
				Fields:    [][2]int{{1, 1}},
				Delimiter: ",",
				Separated: true,
			},
			want: [][]string{{"a"}, {"x"}},
		},
		{
			name: "without separated keeps line without delimiter",
			data: []string{"a,b,c", "hello"},
			config: cfg.Config{
				Fields:    [][2]int{{1, 1}},
				Delimiter: ",",
				Separated: false,
			},
			want: [][]string{{"a"}, {"hello"}},
		},
		{
			name: "out of bounds field does not break later valid field",
			data: []string{"a,b,c"},
			config: cfg.Config{
				Fields:    [][2]int{{1, 1}, {5, 5}, {2, 2}},
				Delimiter: ",",
				Separated: false,
			},
			want: [][]string{{"a", "b"}},
		}}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			data, _ := Cut(tt.data, tt.config)
			if !reflect.DeepEqual(data, tt.want) {
				t.Fatalf("Cut() = %v want %v", data, tt.want)
			}
		})
	}
}
