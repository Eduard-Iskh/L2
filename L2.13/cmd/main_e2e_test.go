package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMainE2E(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		stdin   string
		wantOut string
		wantErr bool
	}{
		{
			name:    "single line with default delimiter",
			args:    []string{"-f", "1"},
			stdin:   "a\tb\tc\n",
			wantOut: "a\n",
			wantErr: false,
		},
		{
			name:    "single line with multiple fields",
			args:    []string{"-f", "1,3", "-d", ","},
			stdin:   "a,b,c\n",
			wantOut: "a,c\n",
			wantErr: false,
		},
		{
			name:    "single line with range",
			args:    []string{"-f", "2-3", "-d", ","},
			stdin:   "a,b,c,d\n",
			wantOut: "b,c\n",
			wantErr: false,
		},
		{
			name:    "single line with out of bounds field",
			args:    []string{"-f", "1,5", "-d", ","},
			stdin:   "a,b,c\n",
			wantOut: "a\n",
			wantErr: false,
		},
		{
			name:    "separated false keeps line without delimiter",
			args:    []string{"-f", "1", "-d", ","},
			stdin:   "hello\n",
			wantOut: "hello\n",
			wantErr: false,
		},
		{
			name:    "separated true skips line without delimiter",
			args:    []string{"-f", "1", "-d", ",", "-s"},
			stdin:   "hello\n",
			wantOut: "",
			wantErr: false,
		},
		{
			name:    "separated true keeps only separated lines",
			args:    []string{"-f", "1", "-d", ",", "-s"},
			stdin:   "a,b,c\nhello\nx,y\n",
			wantOut: "a\nx\n",
			wantErr: false,
		},
		{
			name:    "multiple lines with default delimiter",
			args:    []string{"-f", "2"},
			stdin:   "a\tb\tc\nx\ty\tz\n",
			wantOut: "b\ny\n",
			wantErr: false,
		},
		{
			name:    "multiple lines with custom delimiter",
			args:    []string{"-f", "1,3", "-d", ","},
			stdin:   "a,b,c\nx,y,z\n",
			wantOut: "a,c\nx,z\n",
			wantErr: false,
		},
		{
			name:    "missing required flag f",
			args:    []string{},
			stdin:   "a,b,c\n",
			wantOut: "не указан обязательный флаг -f\n",
			wantErr: true,
		},
		{
			name:    "invalid field format",
			args:    []string{"-f", "1,a", "-d", ","},
			stdin:   "a,b,c\n",
			wantOut: "ошибка формирования конфигурации: неверный элемент списка полей \"a\": ожидается число или диапазон вида N-M\n",
			wantErr: true,
		},
		{
			name:    "invalid reversed range",
			args:    []string{"-f", "5-3", "-d", ","},
			stdin:   "a,b,c\n",
			wantOut: "ошибка формирования конфигурации: неверный диапазон \"5-3\": левая граница больше правой\n",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			args := append([]string{"run", "."}, tt.args...)

			out, err := runGoCommand(t, tt.stdin, args...)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil; output: %q", out)
				}
				if !strings.Contains(out, tt.wantOut) {
					t.Fatalf("error output mismatch:\n got: %q\nwant to contain: %q", out, tt.wantOut)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v; output: %q", err, out)
			}

			if out != tt.wantOut {
				t.Fatalf("output mismatch:\n got: %q\nwant: %q", out, tt.wantOut)
			}
		})
	}
}

func runGoCommand(t *testing.T, stdin string, args ...string) (string, error) {
	t.Helper()

	cmd := exec.Command("go", args...)
	cmd.Dir = "."
	cmd.Env = append(os.Environ(), "GOCACHE=/private/tmp/go-build-cache")

	cmd.Stdin = strings.NewReader(stdin)

	var out bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	return out.String(), err
}
