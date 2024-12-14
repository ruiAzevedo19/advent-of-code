package main

import (
	"fmt"
	"log"
	"os"

	"ceres-search/solution"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("must specify a file path")
	}
	filePath := os.Args[1]

	result, err := solution.CeresSearch(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of XMAS words: %d\n", result.NumberOfWords)
	fmt.Printf("Number of X-MAS words: %d\n", result.NumberOfXWords)
}
