package sorted

import (
	cfg "l210-sort/internal/config"
	"l210-sort/internal/parse"
	"sort"
)

func SizeS(data []cfg.LineComp) []cfg.LineComp {
	sort.Slice(data, func(i, j int) bool {
		a, first := parse.ParseSuffix(data[i].CompElemS)
		b, second := parse.ParseSuffix(data[j].CompElemS)
		if first && second {
			return a < b
		}
		if first {
			return true
		}
		if second {
			return false
		}
		return data[i].CompElemS < data[j].CompElemS
	})
	return data
}
