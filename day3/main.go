package main

import (
	"fmt"
	"log"
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
	point := 0
	for point < len([]rune(input)) {
		inputUpToMul := input[point:]
        nextMul := strings.Index(inputUpToMul, "mul(")
        if nextMul == -1 {
            break
        }
		point += nextMul + MUL_OFFSET

        numbers := []int{}
		numericString := ""
		for {
			asciiValue := int(input[point])
			asciiText := string(rune(asciiValue))
			point++
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
            fmt.Printf("mul(%d,%d) = %d\n", numbers[0], numbers[1], mulResult)
            solution += mulResult
        }
	}
    fmt.Printf("PART I solution: %d", solution)
}
