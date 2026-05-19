package sorted

import (
	cfg "l210-sort/internal/config"
	"strings"
)

func Blank(data []cfg.LineComp) []cfg.LineComp {
	var result []cfg.LineComp
	result = make([]cfg.LineComp, 0, len(data))
	for _, value := range data {
		value.CompElemS = strings.TrimLeft(value.CompElemS, " \t")
		result = append(result, value)
	}
	return result
}
