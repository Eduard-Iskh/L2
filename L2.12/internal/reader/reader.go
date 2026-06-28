package reader

import (
	"bufio"
	"io"
)

// Read читает все строки из указанного файла и возвращает их в формате []string
func Read(fileName io.Reader) ([]string, error) {
	var line string

	scanner := bufio.NewScanner(fileName)

	data := make([]string, 0)
	for scanner.Scan() {
		line = scanner.Text()
		data = append(data, line)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return data, nil
}
