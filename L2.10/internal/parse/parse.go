package parse

import (
	cfg "l210-sort/internal/config"
	"math"
	"strconv"
	"strings"
)

var suffix = map[string]float64{
	"K": math.Pow(2, 10),
	"M": math.Pow(2, 20),
	"G": math.Pow(2, 30),
	"T": math.Pow(2, 40),
}

// find выделяет из каждой строки нужный столбец.
//
// k - номер столбца.
//
// Если k == 0:
// используется вся строка.
//
// Результат сохраняется в CompElemS.
func Find(data [][]byte, k int) []cfg.LineComp {

	dataPars := make([]cfg.LineComp, 0, len(data))

	for _, rawLine := range data {
		line := string(rawLine)

		key := line
		// Если указан номер столбца
		if k != 0 {
			colums := strings.Split(line, "\t")
			if k-1 < len(colums) {
				key = colums[k-1]
			} else {
				key = ""
			}
		}
		dataPars = append(dataPars,
			cfg.LineComp{
				Line:      line,
				CompElemS: key,
			},
		)
	}

	return dataPars
}

func ParseFloat(data []cfg.LineComp) []cfg.LineComp {
	var flag bool
	result := make([]cfg.LineComp, 0, len(data))
	for _, value := range data {
		flag = true
		number, err := strconv.ParseFloat(value.CompElemS, 64)
		if err != nil {
			flag = false
		}
		result = append(result, cfg.LineComp{
			Line:       value.Line,
			CompElemS:  value.CompElemS,
			FCompElemS: number,
			IsNumber:   flag,
		})
	}
	return result
}

func ParseSuffix(data string) (float64, bool) {

	data = strings.TrimSpace(data)
	if len(data) == 0 {
		return 0, false
	}
	last := string(data[len(data)-1])
	value, ok := suffix[last]
	if !ok {
		return 0, false
	}
	number, err := strconv.ParseFloat(data[:len(data)-1], 64)
	if err != nil {
		return 0, false
	}
	return number * value, true
}
