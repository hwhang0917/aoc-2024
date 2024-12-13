package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/hwhang0917/aoc-2024/pkg/utils"
)

var END_PARENTHESIS_ASCII = 41
var COMMA_ASCII = 44
var MUL_OFFSET = 4
var DONT_OFFSET = 7
var DO_OFFSET = 4

func isNumericAscii(asciiValue int) bool {
	return 48 <= asciiValue && asciiValue <= 57
}

func main() {
	input := utils.ReadDayThreeInput(false)
	solution := 0
	pointer := 0
	for pointer < len([]rune(input)) {
		inputUpToMul := input[pointer:]
		nextMul := strings.Index(inputUpToMul, "mul(")
		if nextMul == -1 {
			break
		}
		pointer += nextMul + MUL_OFFSET

		numbers := []int{}
		numericString := ""
		for {
			asciiValue := int(input[pointer])
			asciiText := string(rune(asciiValue))
			pointer++
			if asciiText == ")" {
				break
			}
			if isNumericAscii(asciiValue) {
				numericString = strings.Join([]string{numericString, asciiText}, "")
			} else if asciiValue == COMMA_ASCII {
				num, err := strconv.Atoi(numericString)
				if err != nil {
					log.Fatal("Not a numeric string")
					break
				}
				numbers = append(numbers, num)
				numericString = ""
			} else if asciiValue == END_PARENTHESIS_ASCII {
				break
			} else {
				numericString = ""
				break
			}
		}
		if numericString != "" {
			num, err := strconv.Atoi(numericString)
			if err != nil {
				log.Fatal("Not a numeric string")
				break
			}
			numbers = append(numbers, num)
			numericString = ""
		}
		if len(numbers) == 2 {
			mulResult := numbers[0] * numbers[1]
			solution += mulResult
		}
	}
	fmt.Printf("PART I solution: %d\n", solution)

	type operation struct {
		Name     string
		Position int
	}

	// Reset for Part II
	solution = 0
	r := regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
	matches := r.FindAllStringSubmatch(input, -1)
	enabled := true
	for _, match := range matches {
		switch match[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				first, err := strconv.Atoi(match[2])
				if err != nil {
					log.Fatal("Not a numeric string")
					break
				}
				second, err := strconv.Atoi(match[3])
				if err != nil {
					log.Fatal("Not a numeric string")
					break
				}
				solution += first * second
			}
		}
	}
	fmt.Printf("PART II solution: %d\n", solution)
}
