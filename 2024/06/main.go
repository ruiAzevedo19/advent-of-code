package main

import (
	"fmt"
	"guard-gallivant/solution"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("must specify a file path")
	}
	filePath := os.Args[1]

	result, err := solution.GuardGallivant(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of distinct locations: %d\n", result.GuardDistinctLocations)
}
