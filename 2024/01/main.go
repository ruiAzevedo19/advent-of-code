package main

import (
	"fmt"
	"historian-hysteria/solution"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("must specify a file path")
	}
	filePath := os.Args[1]

	result, err := solution.HistorianHysteria(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("total distance: %d\n", result.TotalDistance)
	fmt.Printf("similarity score: %d\n", result.SimilarityScore)
}
