package main

import (
	"fmt"
	cfg "l210-sort/internal/config"
	"l210-sort/internal/reader"
	"l210-sort/internal/sorted"
	"os"
)

func main() {

	var filename string

	fmt.Println("Введите название файла")

	filename = "L2.10"

	// data1 := "Golang\tfirst\tline\t2\n" +
	// 	"Fasecond\tFeb\t2\t0\n" +
	// 	"Aline\tJan\tthird\t15\n" +
	// 	"Aline\t4r\tad\t7\n" +
	// 	"Zline\tMar\ttext\t100\n" +
	// 	"Bline\tDec\talpha\t1"
	data1 := "apple\t1\n" +
		"banana\t2\n" +
		"orange\t10"
	err := os.WriteFile(filename, []byte(data1), 0644)
	if err != nil {
		fmt.Println("Ошибка записи файла:", err)
		return
	}

	data := reader.Read(filename)
	if len(data) == 0 {
		fmt.Println("Пустой файл на входе")
		return
	}
	config := cfg.NewConfig()

	fmt.Println("\n===== CONFIG =====")
	fmt.Printf("%+v\n", config)

	fmt.Println("\n===== RAW DATA =====")
	for i, line := range data {
		fmt.Printf("[%d] %s\n", i, string(line))
	}

	result := sorted.Sort(config, data)

	fmt.Println("\n===== SORT RESULT =====")

	for i, item := range result {

		fmt.Printf(
			"[%d]\n"+
				"  Line       : %q\n"+
				"  CompElemS  : %q\n"+
				"  FCompElemS : %f\n"+
				"  IsNumber   : %v\n\n",
			i,
			item.Line,
			item.CompElemS,
			item.FCompElemS,
			item.IsNumber,
		)
	}
}
