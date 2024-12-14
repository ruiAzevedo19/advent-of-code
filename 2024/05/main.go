package main

import (
	"fmt"
	"log"
	"os"

	"print-queue/solution"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("specify the input file path")
	}
	inputFilePath := os.Args[1]

	result, err := solution.PrintQueue(inputFilePath)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}

	fmt.Printf("middle page sum: %d\n", result.MiddlePageSumValidUpdates)
	fmt.Printf("middle page sum fixed: %d\n", result.MiddlePageSumFixedUpdates)
}
