package validate

import (
	"fmt"
	cfg "l210-sort/internal/config"
)

func Validate(raw []cfg.LineComp, data []cfg.LineComp) bool {
	fmt.Println(raw)
	fmt.Println(data)
	for i := 1; i < len(data); i++ {
		if raw[i].Line != data[i].Line {
			fmt.Println(raw[i].Line, '\n', data[i].Line)
			fmt.Println(raw[i].Line == data[i].Line)
			return false
		}
	}
	return true
}
