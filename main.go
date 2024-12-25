package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"advent-of-code/challenge"
	_ "advent-of-code/register"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Wrong number of arguments. Specify a year, day and the input file")
	}

	year := os.Args[1]
	day := os.Args[2]
	inputFilePath := os.Args[3]

	solution, err := challenge.SolutionForYearDay(year, day)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	result, err := solution(inputFilePath)
	end := time.Since(start)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	fmt.Printf("Execution time: %fs\n", end.Seconds())
}
