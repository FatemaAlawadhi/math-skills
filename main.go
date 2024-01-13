package main

import (
	"fmt"
	"math-skills/pkg"
	"os"
)

func main() {
	input := os.Args
	if len(input) != 2 {
		fmt.Print("Use: go run main.go data.txt")
	} else {
		Values, err := pkg.FileReader(input[1])
		if err != nil {
			return
		}
		AverageFloat, AverageInt := pkg.Average(Values)
		VarianceFloat, VarianceInt := pkg.Variance(AverageFloat, Values)
		StandardDeviation := pkg.StandardDeviation(VarianceFloat)
		Median := pkg.Median(Values)

		fmt.Printf("Average: %v\n", AverageInt)
		fmt.Printf("Median: %v\n", Median)
		fmt.Printf("Variance: %v\n", VarianceInt)
		fmt.Printf("Standard Deviation: %v\n", StandardDeviation)
	}
}
