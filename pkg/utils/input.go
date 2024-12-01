package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type DayOneInput struct {
    LeftColumn []int
    RightColumn []int
    Lines int
}

var SEPARATOR = "   "

func ReadDayOneInput(testing bool) DayOneInput {
    filename := "input/day1.txt"
    if testing {
        filename = "input/day1_test.txt"
    }
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal("Cannot read %s: %v", filename, err)
    }
    defer file.Close()

    var dayOneInput = DayOneInput{
        LeftColumn: []int{},
        RightColumn: []int{},
        Lines: 0,
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), SEPARATOR)
        leftStr, rightStr := line[0], line[1]
        left, err := strconv.Atoi(leftStr)
        if err != nil {
            log.Fatal("Cannot convert %s to int: %v", leftStr, err)
        }
        right, err := strconv.Atoi(rightStr)
        if err != nil {
            log.Fatal("Cannot convert %s to int: %v", rightStr, err)
        }
        dayOneInput.LeftColumn = append(dayOneInput.LeftColumn, left)
        dayOneInput.RightColumn = append(dayOneInput.RightColumn, right)
        dayOneInput.Lines++
    }
    if err := scanner.Err(); err != nil {
        log.Fatal("Cannot scan %s: %v", filename, err)
    }
    return dayOneInput
}
