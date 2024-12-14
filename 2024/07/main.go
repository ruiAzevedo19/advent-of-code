package main

import (
	"bridge-repair/solution"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("must specify a file path")
	}
	filePath := os.Args[1]

	result, err := solution.BridgeRepair(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total calibration result: %d\n", result.TotalCalibrationResult)
}
