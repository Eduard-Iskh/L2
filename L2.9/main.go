package main

import (
	"fmt"
	"slices"
	"strings"
	"unicode"
)

// Unpacking : реализация примитивной распаковки строки.
// Получает string, возвращает распакованную сроку. Корректно обрабатывает ошибки
func Unpacking(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	runes := []rune(str)
	if unicode.IsDigit(runes[0]) {
		return "", fmt.Errorf("Некоректная строка - начинается с цифры")
	}
	if slices.Equal(runes, []rune{'\\'}) {
		return "", fmt.Errorf("Некоректная строка - только символ экранирования")
	}
	if runes[len(runes)-1] == '\\' {
		return "", fmt.Errorf("Некоректная строка - последий символ не должен быть символом экранирования")
	}
	var unpack string = ""
	prev := string(runes[0])
	var flag bool = false

	for i := 0; i < len(runes); i++ {
		ch := runes[i]
		if ch == '\\' {
			flag = true
			continue
		}
		if flag {
			flag = false
			prev = string(ch)
			unpack += string(ch)
			continue
		}

		if unicode.IsDigit(ch) {
			if unicode.IsDigit(runes[i-1]) && runes[i-2] != '\\' {
				return "", fmt.Errorf("Некоректная строка - две цифры подряд без экранирования")
			}
			unpack += strings.Repeat(prev, int(ch-'0'))
			unpack = unpack[:len(unpack)-1]
			continue
		}
		unpack += string(ch)
		prev = string(ch)
	}
	return unpack, nil
}

func main() {
	var str string
	fmt.Println("Введите строку")
	fmt.Scan(&str)
	str = "a\\0\\"
	ans, err := Unpacking(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ans)
}
