package sorted

import (
	"fmt"
	cfg "l210-sort/internal/config"
	"l210-sort/internal/dedup"
	"l210-sort/internal/parse"
	"l210-sort/internal/validate"
	"sort"
)

// sorted сортирует массив структур LineComp.
func Sorted(data []cfg.LineComp, flag ...bool) []cfg.LineComp {

	var sortedStruct []cfg.LineComp = data
	sort.Slice(sortedStruct, func(i, j int) bool {
		return sortedStruct[i].CompElemS <
			sortedStruct[j].CompElemS
	})
	return sortedStruct
}

// Sort функция
func Sort(config cfg.Config, data [][]byte) []cfg.LineComp {
	newMap := parse.Find(data, config.K)
	raw := newMap
	fmt.Println("\n", "raw = ", raw[0].Line, "\n", "end")
	if config.B {
		newMap = Blank(newMap)
	}
	newMap = Sorted(newMap)
	if config.N {
		newMap = Numbers(parse.ParseFloat(newMap))
	}

	if config.U {
		newMap = dedup.Unic(newMap)
	}

	if config.M {
		newMap = MonthS(newMap)
	}

	if config.R {
		Reverse(newMap)
	}

	if config.C {
		if !validate.Validate(raw, newMap) {
			fmt.Println("Данные не отсортированы")
		}
	}
	if config.H {
		newMap = SizeS(newMap)
	}
	return newMap
}
