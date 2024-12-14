package main

import (
	"fmt"
	"log"
	"os"
	"red-nosed-reports/solution"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("must specify a file path")
	}
	filePath := os.Args[1]

	result, err := solution.RedNosedReports(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Numer of correct reports: %d\n", result.NumberOfCorrectReports)
	fmt.Printf("Numer of fixed reports: %d\n", result.NumberOfFixedReports)
	fmt.Printf("Total: %d\n", result.NumberOfCorrectReports+result.NumberOfFixedReports)
}
