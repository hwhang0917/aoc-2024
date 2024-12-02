package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/hwhang0917/aoc-2024/pkg/utils"
)

func isSafe(report []int) bool {
	isIncreasing := report[1] > report[0]
	for i := 1; i < len(report); i++ {
		if report[i] == report[i-1] {
			return false
		}
		if isIncreasing != (report[i] > report[i-1]) {
			return false
		}
		if math.Abs(float64(report[i]-report[i-1])) > 3 {
			return false
		}
	}
	return true
}

func problemDampener(report []int, idx int) []int {
    return slices.Concat(report[:idx], report[idx+1:])
}

func main() {
	input := utils.ReadDayTwoInput(true)

	// Part I
	safeCount := 0
	for _, report := range input.Reports {
		if isSafe(report) {
			safeCount++
		}
	}
	fmt.Printf("PART I solution: %d\n", safeCount)

	// Part II
	safeCount = 0
	for _, report := range input.Reports {
		for idx := range report {
            reportSlice := problemDampener(report, idx)
			if isSafe(reportSlice) {
				safeCount++
				break
			}
		}
	}
	fmt.Printf("PART II solution: %d\n", safeCount)
}
