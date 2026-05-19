package sorted

import (
	cfg "l210-sort/internal/config"
	"sort"
)

var month = map[string]int{
	"Jan": 1,
	"Feb": 2,
	"Mar": 3,
	"Apr": 4,
	"May": 5,
	"Jun": 6,
	"Jul": 7,
	"Aug": 8,
	"Sep": 9,
	"Oct": 10,
	"Nov": 11,
	"Dec": 12,
}

func MonthS(data []cfg.LineComp) []cfg.LineComp {

	var SortedSlice []cfg.LineComp = data

	sort.Slice(SortedSlice, func(i, j int) bool {
		return month[SortedSlice[i].CompElemS] < month[SortedSlice[j].CompElemS]
	})

	return SortedSlice

}
