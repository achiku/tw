package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var input = flag.String("file", "", "input file name")

func main() {
	flag.Parse()
	if *input == "" {
		log.Fatal("-file must be specified")
	}
	f, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	sts, err := parser(f)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range sts {
		fmt.Printf("%s\n", format(s))
	}
}
