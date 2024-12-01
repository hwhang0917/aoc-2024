package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/hwhang0917/aoc-2024/pkg/utils"
)

func main() {
    input := utils.ReadDayOneInput(false)

    // Part 1
    sortedLeftColumn := make([]int, len(input.LeftColumn))
    sortedRightColumn := make([]int, len(input.RightColumn))
    copy(sortedLeftColumn, input.LeftColumn)
    copy(sortedRightColumn, input.RightColumn)
    sort.Ints(sortedLeftColumn)
    sort.Ints(sortedRightColumn)

    distanceSum := 0

    for i := 0; i < input.Lines; i++ {
        left := sortedLeftColumn[i]
        right := sortedRightColumn[i]

        distance := math.Abs(float64(left - right))
        distanceSum += int(distance)
    }

    fmt.Printf("PART I solution: %d\n", distanceSum)

    // Part 2
    similarSum := 0
    frequencyMap := make(map[int]int)
    for _, num := range input.RightColumn {
        frequencyMap[num]++
    }
    for i := 0; i < input.Lines; i++ {
        left := input.LeftColumn[i]
        similarity := frequencyMap[left]
        similarSum += similarity * left
    }

    fmt.Printf("PART II solution: %d\n", similarSum)
}

