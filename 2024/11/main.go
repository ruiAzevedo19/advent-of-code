package main

import (
	"fmt"
	"log"
	"os"

	"plutonian-pebbles/solution"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("must specify a file path")
	}
	filePath := os.Args[1]

	result, err := solution.PlutonianPebbles(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of stones (25 blinks): %d\n", result.NumberOfStones25)
	fmt.Printf("Number of stones (75 blinks): %d\n", result.NumberOfStones75)
}
