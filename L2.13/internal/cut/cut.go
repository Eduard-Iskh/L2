package cut

import (
	"strings"

	"example.com/internal/cfg"
)

func parseFields(line []string, fields [][2]int) []string {
	var start, end int
	data := make([]string, 0)
	for _, nums := range fields {
		start = nums[0]
		end = nums[1]
		for i := start - 1; i < end; i++ {
			if i < len(line) {
				data = append(data, line[i])
			} else {
				break
			}
		}
	}
	return data
}

// Cut выделяет указанные поля из входных строк по заданному разделителю
func Cut(data []string, config cfg.Config) ([][]string, error) {
	var line []string
	field := make([][]string, 0)
	for _, i := range data {
		line = strings.Split(i, config.Delimiter)
		if config.Separated {
			if len(line) <= 1 {
				continue
			}
		}
		field = append(field, parseFields(line, config.Fields))
	}

	return field, nil
}
