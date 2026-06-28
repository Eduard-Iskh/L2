package cfg

import "flag"

// Config хранит итоговую конфигурацию для запуска grep
type Config struct {
	Pattern    string
	After      int  //-A
	Before     int  //-B
	Context    int  //-C
	Count      bool //-c
	IgnoreCase bool //-i
	Invert     bool //-v
	Fixed      bool //-F
	Num        bool //-n
}

// Flags хранит указатели на зарегистрированные флаги командной строки
type Flags struct {
	After      *int  //-A
	Before     *int  //-B
	Context    *int  //-C
	Count      *bool //-c
	IgnoreCase *bool //-i
	Invert     *bool //-v
	Fixed      *bool //-F
	Num        *bool //-n
}

// NewFlag регистрирует флаги командной строки и возвращает их ссылки
func NewFlag() Flags {
	return Flags{
		After:      flag.Int("A", 0, "Вывод N строк после каждой найденной строки"),
		Before:     flag.Int("B", 0, "Вывод N строк до каждой найденной строки"),
		Context:    flag.Int("C", 0, "Вывод N строк контекста вокруг найденной строки"),
		Count:      flag.Bool("c", false, "Выводить количество совпадающих строк"),
		IgnoreCase: flag.Bool("i", false, "Игнорировать регистр"),
		Invert:     flag.Bool("v", false, "Инвертировать фильтр"),
		Fixed:      flag.Bool("F", false, "Шаблон как фиксированная строка"),
		Num:        flag.Bool("n", false, "Номер строки перед каждой найденной строкой"),
	}
}

// NewConfig собирает итоговую конфигурацию из шаблона и значений флагов
func NewConfig(pattern string, flags Flags) Config {

	return Config{
		Pattern:    pattern,
		After:      *flags.After,
		Before:     *flags.Before,
		Context:    *flags.Context,
		Count:      *flags.Count,
		IgnoreCase: *flags.IgnoreCase,
		Invert:     *flags.Invert,
		Fixed:      *flags.Fixed,
		Num:        *flags.Num,
	}
}
