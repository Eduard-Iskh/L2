package main

import (
	"flag"
	"fmt"
	"os"

	"example.com/internal/cfg"
	"example.com/internal/cut"
	"example.com/internal/printer"
	"example.com/internal/reader"
)

func main() {
	flags := cfg.NewFlag()
	flag.Parse()
	config, err := cfg.NewConfig(flags)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data, err := reader.Read(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fields, err := cut.Cut(data, config)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printer.PrintAll(fields, config)
}
