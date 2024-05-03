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

    var sum int
    for _, line := range lines {
        re := regexp.MustCompile("[0-9]")
        numString := re.FindAllString(line, -1)
        if numString != nil {
            fmt.Println(numString)
            c, _ := getNumber(numString)
            sum += c
            fmt.Printf("check c: %v, sum: %v \n", c, sum)
        }
    }
}

//func sum() func(line string) int {
//    var sum int
//    return func(line string) int {
//        re := regexp.MustCompile("[0-9]")
//        numString := re.FindAllString(line, -1)
//        if numString != nil {
//            c, _ := getNumber(numString)
//            sum += c
//        }
//
//        return sum
//    }
//}

func getNumber(numString []string) (int, error) {
    calibrationValue := string(numString[0])+string(numString[len(numString)-1])

    return strconv.Atoi(calibrationValue)
}

func splitText(text string, sep string) []string{
    return strings.Split(text, sep)
}

