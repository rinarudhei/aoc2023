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

    s := sum()
    for _, line := range lines {
        s(line)
    }
}

func sum() func(line string) {
    var total int
    re := regexp.MustCompile("[0-9]")
    return func(line string) {
        numString := re.FindAllString(line, -1)
        if numString != nil {
            c, _ := getNumber(numString)
            total += c
        }
        fmt.Printf("check total: %v \n", total)
    }
}

func getNumber(numString []string) (int, error) {
    calibrationValue := string(numString[0])+string(numString[len(numString)-1])

    return strconv.Atoi(calibrationValue)
}

func splitText(text string, sep string) []string{
    return strings.Split(text, sep)
}

