// https://adventofcode.com/2023/day/1

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
    text, err := os.ReadFile("day1_in.txt")
    if err != nil {
        fmt.Printf("error: %v \n", err)
    }

    lines := splitText(string(text), "\n")

    s := Sum()
    for _, line := range lines {
        s(line)
    }
}

func Sum() func(line string) int {
    var total int
    reFirst := regexp.MustCompile("([0-9]|one|two|three|four|five|six|seven|eight|nine)")
    reLast := regexp.MustCompile("([0-9]|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)")
    return func(line string) int {
        first := reFirst.FindString(line)
        last := reLast.FindString(reverse(line))
        if first != "" {
            firstDigit := convertToDigit(first)
            lastDigit := convertToDigit(last)
            val, _ := strconv.Atoi(firstDigit+lastDigit)
            total += val
            fmt.Printf("check total: %v \n", total)
        }

        return total
    }
}

func reverse(line string) string {
    s := []rune(line)
    for i, j := 0, len(line)-1; i < len(line)/2; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }

    return string(s)
}


func convertToDigit(numString string) string {
    switch {
    case numString == "one" || numString == "eno" :
        return "1"
    case numString == "two" || numString == "owt":
        return "2"
    case numString == "three" || numString == "eerht":
        return "3"
    case numString == "four" || numString == "ruof":
        return "4"
    case numString == "five" || numString == "evif":
        return "5"
    case numString == "six" || numString == "xis":
        return "6"
    case numString == "seven" || numString == "neves":
        return "7"
    case numString == "eight" || numString == "thgie":
        return "8"
    case numString == "nine" || numString == "enin":
        return "9"
    default:
        return numString
    }
}

func splitText(text string, sep string) []string{
    return strings.Split(text, sep)
}

