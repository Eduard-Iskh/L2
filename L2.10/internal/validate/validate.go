package validate

import (
	"l210-sort/internal/compare"
	cfg "l210-sort/internal/config"
)

func Validate(data []cfg.LineComp, conf cfg.Config) bool {

	for i := 1; i < len(data); i++ {
		if compare.Less(data[i], data[i-1], conf) {
			return false
		}
	}
	return true
}
