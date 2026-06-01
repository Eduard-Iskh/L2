package sorted

import (
	cfg "l210-sort/internal/config"
	"sort"
)

var month = map[string]int{
	"jan": 1,
	"feb": 2,
	"mar": 3,
	"apr": 4,
	"may": 5,
	"jun": 6,
	"jul": 7,
	"aug": 8,
	"sep": 9,
	"oct": 10,
	"nov": 11,
	"dec": 12,
}

func MonthS(data []cfg.LineComp) []cfg.LineComp {

	var SortedSlice []cfg.LineComp = data

	sort.Slice(SortedSlice, func(i, j int) bool {
		return month[SortedSlice[i].CompElemS] < month[SortedSlice[j].CompElemS]
	})

	return SortedSlice

}
