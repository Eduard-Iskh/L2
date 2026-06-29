package printer

import (
	"fmt"
	"strings"

	"example.com/internal/cfg"
)

// PrintAll функция вывода
func PrintAll(data [][]string, config cfg.Config) {
	for _, line := range data {
		fmt.Println(strings.Join(line, config.Delimiter))
	}
}
