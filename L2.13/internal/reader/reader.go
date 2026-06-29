package reader

import (
	"bufio"
	"io"
)

// Read функция чтения входных данных
func Read(stdin io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(stdin)

	data := make([]string, 0)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return data, nil
}
