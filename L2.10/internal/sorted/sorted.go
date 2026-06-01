package sorted

import (
	"fmt"
	cfg "l210-sort/internal/config"
	"l210-sort/internal/parse"
	"l210-sort/validate"
	"sort"
)

// sorted сортирует массив структур LineComp.
func Sorted(data []cfg.LineComp, config cfg.Config) []cfg.LineComp {

	var sortedStruct []cfg.LineComp = data
	sort.Slice(sortedStruct, func(i, j int) bool {
		return Less(sortedStruct[i], sortedStruct[j], config)
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
	newMap = Sorted(newMap, config)

	if config.C {
		if !validate.Validate(newMap) {
			fmt.Println("Данные не отсортированы")
		}
	}

	return newMap
}
