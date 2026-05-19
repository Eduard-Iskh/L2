package reader

import (
	"bufio"
	"fmt"
	"os"
)

func Read(filename string) [][]byte {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Bytes())
	}
	return data
}
