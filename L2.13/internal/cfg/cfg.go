package cfg

import (
	"errors"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	reFirst  = regexp.MustCompile("^[1-9][0-9]*$")
	reSecond = regexp.MustCompile("^[1-9][0-9]*-[1-9][0-9]*$")
)

// Config хранит итоговую конфигурацию для запуска cut
type Config struct {
	Fields    [][2]int //-f
	Delimiter string   //-d
	Separated bool     //-s
}

// Flag хранит указатели на зарегестрированные флаги командной строки
type Flag struct {
	Fields    *string
	Delimiter *string
	Separated *bool
}

// NewFlag регистрирует флаги командной строки и возвращает их
func NewFlag() Flag {
	return Flag{
		Fields:    flag.String("f", "", ""),
		Delimiter: flag.String("d", "\t", ""),
		Separated: flag.Bool("s", false, ""),
	}
}

func parseFields(fields string) ([][2]int, error) {
	parts := strings.Split(fields, ",")
	data := make([][2]int, 0)
	var start, end int
	for _, i := range parts {
		if reFirst.MatchString(i) {
			start, _ = strconv.Atoi(i)
			end = start
		} else if reSecond.MatchString(i) {
			nums := strings.Split(i, "-")
			start, _ = strconv.Atoi(nums[0])
			end, _ = strconv.Atoi(nums[1])
			if start > end {
				return nil, fmt.Errorf("неверный диапазон %q: левая граница больше правой", i)
			}

		} else {
			return nil, fmt.Errorf("неверный элемент списка полей %q: ожидается число или диапазон вида N-M", i)
		}
		data = append(data, [2]int{start, end})
	}
	return data, nil
}

// NewConfig собирает итоговую конфигурацию
func NewConfig(flags Flag) (Config, error) {
	var err error
	var data [][2]int
	if *flags.Fields == "" {
		return Config{}, errors.New("не указан обязательный флаг -f")
	}
	data, err = parseFields(*flags.Fields)

	if err != nil {
		return Config{}, fmt.Errorf("ошибка формирования конфигурации: %w", err)
	}
	return Config{
		Fields:    data,
		Delimiter: *flags.Delimiter,
		Separated: *flags.Separated,
	}, nil
}
