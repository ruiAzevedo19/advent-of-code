package main

import (
	"fmt"
	"log"
	"mull-it-over/solution"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("must specify a file path")
	}
	filePath := os.Args[1]

	result, err := solution.MullItOver(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Summed multiplications: %d\n", result.SummedMultiplications)
	fmt.Printf("Summed multiplication with do-instructions: %d\n", result.SummedMultiplicationsDo)
}
