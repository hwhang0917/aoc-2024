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


func ReadDayOneInput(testing bool) DayOneInput {
    seperator := "   "
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
        rawLine := scanner.Text()
        if rawLine == "" {
            continue
        }
        line := strings.Split(rawLine, seperator)
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

type DayTwoInput struct {
    Reports [][]int
    Lines int
}

func ReadDayTwoInput(testing bool) DayTwoInput {
    seperator := " "
    filename := "input/day2.txt"
    if testing {
        filename = "input/day2_test.txt"
    }
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal("Cannot read %s: %v", filename, err)
    }
    defer file.Close()

    var dayTwoInput = DayTwoInput{
        Reports: [][]int{},
        Lines: 0,
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        rawLine := scanner.Text()
        if rawLine == "" {
            continue
        }
        lineString := strings.Split(rawLine, seperator)
        var line []int
        for _, numStr := range lineString {
            num, err := strconv.Atoi(numStr)
            if err != nil {
                log.Fatal("Cannot convert %s to int: %v", numStr, err)
            }
            line = append(line, num)
        }
        dayTwoInput.Reports = append(dayTwoInput.Reports, line)
        dayTwoInput.Lines++
    }
    return dayTwoInput
}
