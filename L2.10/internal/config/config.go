package config

import "flag"

// Config хранит значения всех флагов командной строки.
//
// k - номер столбца для сортировки
// n - сортировка как чисел
// r - обратный порядок сортировки
// u - вывод только уникальных элементов
type Config struct {
	K int
	N bool
	R bool
	U bool
	M bool
	B bool
	C bool
	H bool
}

// LineComp хранит строку и значение,
// по которому будет происходить сортировка.
//
// Line       - исходная строка
// CompElemS  - строковое значение для сравнения
// FCompElemS - числовое значение для сравнения
type LineComp struct {
	Line       string
	CompElemS  string
	FCompElemS float64
	IsNumber   bool
}

// NewConfig читает флаги командной строки
// и сохраняет их в структуру Config.
//
// Пример запуска:
// go run main.go -k=2 -n -r
func NewConfig() Config {
	k := flag.Int("k", 0, "номер столбца для сортировки")
	n := flag.Bool("n", false, "сортировка по числовому значению")
	r := flag.Bool("r", false, "сортировка в обратном порядке")
	u := flag.Bool("u", false, "вывод только уникальные элементы")
	m := flag.Bool("M", false, "сортировать по названию месяца ")
	b := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	c := flag.Bool("c", false, "проверка сортировки")
	h := flag.Bool("h", false, "сортировать по числовому значению")
	// Считываем флаги из командной строки
	flag.Parse()

	return Config{
		K: *k,
		N: *n,
		R: *r,
		U: *u,
		M: *m,
		B: *b,
		C: *c,
		H: *h,
	}
}
