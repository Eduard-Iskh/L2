package dedup

import (
	cfg "l210-sort/internal/config"
)

// unic удаляет повторяющиеся элементы.
//
// Уникальность определяется
// по полю CompElemS.
//
// map используется как множество:
// если ключ уже встречался —
// элемент пропускается.

func Unic(data []cfg.LineComp) []cfg.LineComp {
	seen := make(map[string]bool, 0)

	res := make([]cfg.LineComp, 0, len(data))

	for _, value := range data {

		// Если элемент ещё не встречался
		if !seen[value.CompElemS] {

			// Помечаем как встреченный
			seen[value.CompElemS] = true

			// Добавляем в результат
			res = append(res, value)
		}
	}

	return res
}
