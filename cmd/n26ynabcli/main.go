package main

import (
	"os"
	"log"
	"github.com/Dmitry-Erokhin/n26ynab"
	"fmt"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: n26ynabcli <N26 csv dump> <YNAB target file>")
		os.Exit(0)
	}

	input := os.Args[1]
	output := os.Args[2]

	log.Printf("Will create YNAB data in %s from N26 data in %s", output, input)

	in, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	out, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	err = budgeter.ConvertData(in, out)
	if err != nil {
		log.Fatal(err)
	}
}
