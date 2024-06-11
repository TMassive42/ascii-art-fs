package main

import (
	"fmt"
	"log"
	"os"

	"ascii/ascii"
)

// formats is a map that maps output format flags to their corresponding filenames.
var formats = map[string]string{
	"standard":   "standard.txt",
	"thinkertoy": "thinkertoy.txt",
	"shadow":     "shadow.txt",
}

func main() {
	args := os.Args[1:]
	if len(args) == 0  || len(args) > 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}

	input := args[0]
	var banner string

	if len(args) > 1 {
		banner = args[1]
	} else {
		banner = "standard"
	}

	filename, ok := formats[banner]
	if !ok {
		fmt.Println("Invalid banner specified.")
		return
	}

	asciMap, err := ascii.ReadASCIIMapFromFile(filename)
	if err != nil {
		log.Fatalf("Error reading ASCII map: %v", err)
	}

	err = ascii.PrintArt(input, asciMap)
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
