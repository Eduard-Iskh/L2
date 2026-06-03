package dedup

import (
	"l210-sort/internal/compare"
	cfg "l210-sort/internal/config"
)

func Unic(data []cfg.LineComp, conf cfg.Config) []cfg.LineComp {
	res := []cfg.LineComp{data[0]}
	for i := 1; i < len(data); i++ {
		if compare.Compare(data[i-1], data[i], conf) != 0 {
			res = append(res, data[i])
		}
	}
	return res
}
