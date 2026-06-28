package printer

import (
	"fmt"

	"example.com/internal/cfg"
)

// Print выводит результат поиска в зависимости от переданных флагов
func Print(data []string, merge [][2]int, counMatches int, config cfg.Config) {
	if config.Count {
		fmt.Println(counMatches)
		return
	}
	var add string = ""
	for _, i := range merge {
		for j := i[0]; j < i[1]+1; j++ {
			if config.Num {
				add = fmt.Sprintf("%d: ", j+1)
			}
			fmt.Println(add + data[j])
		}
	}
}
