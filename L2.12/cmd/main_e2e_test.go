package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestMainE2E(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		fileContent string
		stdin       string
		wantOut     string
		wantErr     bool
	}{
		{
			name: "basic regexp from file",
			args: []string{"apple"},
			fileContent: strings.Join([]string{
				"apple",
				"banana",
				"pineapple",
			}, "\n"),
			wantOut: "apple\npineapple\n",
		},
		{
			name: "fixed string from file",
			args: []string{"-F", "a.c"},
			fileContent: strings.Join([]string{
				"a.c",
				"abc",
				"x a.c y",
			}, "\n"),
			wantOut: "a.c\nx a.c y\n",
		},
		{
			name: "ignore case and line numbers",
			args: []string{"-i", "-n", "apple"},
			fileContent: strings.Join([]string{
				"Apple",
				"banana",
				"APPLE",
			}, "\n"),
			wantOut: "1: Apple\n3: APPLE\n",
		},
		{
			name: "count matches",
			args: []string{"-c", "apple"},
			fileContent: strings.Join([]string{
				"apple",
				"banana",
				"apple",
			}, "\n"),
			wantOut: "2\n",
		},
		{
			name: "invert filter",
			args: []string{"-v", "apple"},
			fileContent: strings.Join([]string{
				"apple",
				"banana",
				"pear",
			}, "\n"),
			wantOut: "banana\npear\n",
		},
		{
			name: "context around match",
			args: []string{"-C", "1", "banana"},
			fileContent: strings.Join([]string{
				"apple",
				"banana",
				"pear",
			}, "\n"),
			wantOut: "apple\nbanana\npear\n",
		},
		{
			name: "stdin input",
			args: []string{"-i", "apple"},
			stdin: strings.Join([]string{
				"Apple",
				"banana",
			}, "\n"),
			wantOut: "Apple\n",
		},
		{
			name:        "invalid regexp",
			args:        []string{"("},
			fileContent: "apple",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			args := append([]string{"run", "."}, tt.args...)

			if tt.fileContent != "" {
				filePath := writeTempInputFile(t, tt.fileContent)
				args = append(args, filePath)
			}

			out, err := runGoCommand(t, tt.stdin, args...)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil; output: %q", out)
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

	if stdin != "" {
		cmd.Stdin = strings.NewReader(stdin)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	return out.String(), err
}

func writeTempInputFile(t *testing.T, content string) string {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, "input.txt")

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("write temp file: %v", err)
	}

	return path
}
