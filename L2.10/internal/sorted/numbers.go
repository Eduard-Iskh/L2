package sorted

import (
	cfg "l210-sort/internal/config"
	"sort"
)

func Numbers(data []cfg.LineComp) []cfg.LineComp {

	// Сортировка чисел
	sort.Slice(data, func(i, j int) bool {
		a := data[i]
		b := data[j]

		if a.IsNumber && b.IsNumber {
			return a.FCompElemS < b.FCompElemS
		}
		if a.IsNumber {
			return true
		}
		if b.IsNumber {
			return false
		}
		return a.CompElemS < b.CompElemS
	})
	// false -> сортировка по числам
	return data
}
