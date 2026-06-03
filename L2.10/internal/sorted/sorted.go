package sorted

import (
	"fmt"
	"l210-sort/internal/compare"
	cfg "l210-sort/internal/config"
	"l210-sort/internal/dedup"
	"l210-sort/internal/parse"
	"l210-sort/internal/validate"
	"sort"
)

// sorted сортирует массив структур LineComp.
func Sorted(data []cfg.LineComp, config cfg.Config) []cfg.LineComp {

	var sortedStruct []cfg.LineComp = data
	sort.Slice(sortedStruct, func(i, j int) bool {
		return compare.Less(sortedStruct[i], sortedStruct[j], config)
	})
	return sortedStruct
}

// Sort функция
func Sort(config cfg.Config, data [][]byte) []cfg.LineComp {
	newMap := parse.Find(data, config.K)
	raw := newMap
	fmt.Println("\n", "raw = ", raw[0].Line, "\n", "end")
	if config.N {
		newMap = parse.ParseFloat(newMap)
	}
	if config.C {
		if !validate.Validate(newMap, config) {
			fmt.Println("Данные не отсортированы")
		}
	}

	newMap = Sorted(newMap, config)
	if config.U {
		newMap = dedup.Unic(newMap, config)
	}

	return newMap
}
