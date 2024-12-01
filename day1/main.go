package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/hwhang0917/aoc-2024/pkg/utils"
)

func main() {
    input := utils.ReadDayOneInput(false)

    sort.Ints(input.LeftColumn)
    sort.Ints(input.RightColumn)

    distanceSum := 0

    for i := 0; i < input.Lines; i++ {
        left := input.LeftColumn[i]
        right := input.RightColumn[i]

        distance := math.Abs(float64(left - right))
        distanceSum += int(distance)
    }

    fmt.Println(distanceSum)
}

