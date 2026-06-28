package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)

}

func main() {
	var key string
	var groupKey string
	res := make(map[string][]string)
	anagrams := make(map[string][]string)
	words := []string{"пятка", "тяпка", "Пятак", "слиток", "лИСток", "столик", "стол"}
	for _, value := range words {
		value = strings.ToLower(value)
		key = sortString(value)
		anagrams[key] = append(anagrams[key], value)
	}
	for _, value := range anagrams {
		if len(value) > 1 {
			groupKey = value[0]
			sort.Slice(value, func(i, j int) bool {
				return value[i] < value[j]
			})
			res[groupKey] = value
		}
	}
	fmt.Println(res)
}
