package grep

import (
	"regexp"

	"example.com/internal/cfg"
)

// Run функция формирования среза номеров строк, в которых присутствует строки, соответствующие
// заданному шаблону
func Run(data []string, config cfg.Config, isMatch func(line string, config cfg.Config, re *regexp.Regexp) bool) ([][2]int, int, error) {
	var matchIndex []int
	var match bool
	var re *regexp.Regexp
	var err error
	var counMatches int = 0
	if !config.Fixed {
		if config.IgnoreCase {
			re, err = regexp.Compile(`(?i)` + config.Pattern)
		} else {
			re, err = regexp.Compile(config.Pattern)
		}
		if err != nil {
			return nil, 0, err
		}
	}

	for i, line := range data {
		match = isMatch(line, config, re)
		if config.Invert {
			match = !match
		}
		if match {
			matchIndex = append(matchIndex, i)
			counMatches++
		}
	}
	var ranges [][2]int
	var start, end, after, before int
	if config.Context != 0 {
		before = config.Context
		after = config.Context
	} else {
		before = config.Before
		after = config.After
	}
	//создание списка всех промежутков контекста
	for _, i := range matchIndex {
		start = i - before
		end = i + after
		if start < 0 {
			start = 0
		}
		if end > len(data)-1 {
			end = len(data) - 1
		}
		ranges = append(ranges, [2]int{start, end})
	}
	if len(ranges) == 0 {
		return nil, 0, nil
	}
	merge := [][2]int{ranges[0]}
	for _, i := range ranges[1:] {
		last := &merge[len(merge)-1]
		if i[0] <= last[1]+1 {
			if i[1] > last[1] {
				last[1] = i[1]
			}
		} else {
			merge = append(merge, i)
		}
	}
	return merge, counMatches, nil
}
