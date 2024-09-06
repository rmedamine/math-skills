package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"skils/function"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		// err
	}
	file, _ := os.ReadFile(arg[0])
	slice := strings.Split(string(file), "\n")
	x := []float64{}
	for _, item := range slice {
		num, err := strconv.ParseFloat(item, 64)
		if err != nil {
			continue
		}
		x = append(x, num)
	}
	fmt.Println("average", (function.AverageCalculator(x)))
	fmt.Println("median", (function.MedianCalculator(x)))
	fmt.Println("variance", int(function.VarianceCalculator(x)))
	fmt.Println("standarddeviation", (function.StandardDeviation(x)))
}
