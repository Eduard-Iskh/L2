package cfg

import (
	"slices"
	"testing"
)

func TestNewConfig(t *testing.T) {
	test := []struct {
		name       string
		flags      Flag
		wantConfig Config
		wantErr    bool
	}{
		{
			name: "single field",
			flags: Flag{
				Fields:    stringPtr("1"),
				Delimiter: stringPtr("\t"),
				Separated: boolPtr(false),
			},
			wantConfig: Config{
				Fields:    [][2]int{{1, 1}},
				Delimiter: "\t",
				Separated: false,
			},
			wantErr: false,
		},
		{
			name: "single field and range",
			flags: Flag{
				Fields:    stringPtr("1,3-5"),
				Delimiter: stringPtr("\t"),
				Separated: boolPtr(false),
			},
			wantConfig: Config{
				Fields:    [][2]int{{1, 1}, {3, 5}},
				Delimiter: "\t",
				Separated: false,
			},
			wantErr: false,
		},
		{
			name: "custom delimiter and separated flag",
			flags: Flag{
				Fields:    stringPtr("2,4"),
				Delimiter: stringPtr(","),
				Separated: boolPtr(true),
			},
			wantConfig: Config{
				Fields:    [][2]int{{2, 2}, {4, 4}},
				Delimiter: ",",
				Separated: true,
			},
			wantErr: false,
		},
		{
			name: "invalid zero field",
			flags: Flag{
				Fields:    stringPtr("0"),
				Delimiter: stringPtr("\t"),
				Separated: boolPtr(false),
			},
			wantConfig: Config{},
			wantErr:    true,
		},
		{
			name: "invalid format with letters",
			flags: Flag{
				Fields:    stringPtr("1,a"),
				Delimiter: stringPtr("\t"),
				Separated: boolPtr(false),
			},
			wantConfig: Config{},
			wantErr:    true,
		},
		{
			name: "invalid reversed range",
			flags: Flag{
				Fields:    stringPtr("5-3"),
				Delimiter: stringPtr("\t"),
				Separated: boolPtr(false),
			},
			wantConfig: Config{},
			wantErr:    true,
		},
		{
			name: "invalid empty part",
			flags: Flag{
				Fields:    stringPtr("1,,3"),
				Delimiter: stringPtr("\t"),
				Separated: boolPtr(false),
			},
			wantConfig: Config{},
			wantErr:    true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			config, err := NewConfig(tt.flags)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewConfig() error = %v, wantError %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !slices.Equal(config.Fields, tt.wantConfig.Fields) {
				t.Fatalf("NewConfig() неправильный парсинг полей: parse = %v,  want %v", config.Fields, tt.wantConfig.Fields)
			}

			if config.Delimiter != tt.wantConfig.Delimiter {
				t.Fatalf("NewConfig() Delimiter = %v, vant %v", config.Delimiter, tt.wantConfig.Delimiter)
			}
			if config.Separated != tt.wantConfig.Separated {
				t.Fatalf("NewConfig() Separated = %v, vant %v", config.Separated, tt.wantConfig.Separated)
			}
		})
	}

}

func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}
