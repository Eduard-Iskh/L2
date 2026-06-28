package main

import (
	"flag"
	"fmt"
	"os"

	"example.com/internal/cfg"
	"example.com/internal/grep"
	"example.com/internal/match"
	"example.com/internal/printer"
	"example.com/internal/reader"
)

func main() {
	flags := cfg.NewFlag()
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("формат grep:[flags] pattern [file name]")
		os.Exit(1)
	}

	var data []string
	var err error

	if flag.NArg() >= 2 {
		file, err := os.Open(flag.Arg(1))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		data, err = reader.Read(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		data, err = reader.Read(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	pattern := flag.Arg(0)

	config := cfg.NewConfig(pattern, flags)
	if config.Fixed {
	}
	merge, counMatches, err := grep.Run(data, config, match.IsMatch)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printer.Print(data, merge, counMatches, config)
}
