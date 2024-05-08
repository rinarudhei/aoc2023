package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

)

var ins, _ = os.ReadFile("day3_in.txt")
var intemp = strings.Split(string(ins), "\n")
var in = intemp[0:len(intemp)-1]

func main() {
    var sum int
    var sumGears int
    for row, line := range in {
        sum += countSum(line, row)
        sumGears += countSumGear(row)
    }

    fmt.Printf("check sum: %v \n", sum)
    fmt.Printf("check sumGears: %v \n", sumGears)
    fmt.Printf("in lenth: %v \n", len(in))
}

func countSumGear(row int) int {
    var sum int
    var line = in[row]
    reGear := regexp.MustCompile("[*]")
    gearIndexes := reGear.FindAllStringIndex(line, -1)
    for _, i := range gearIndexes {
        isGear, numbers := getGearNumbers(i[0], row)
        if isGear {
            sum += numbers[0]*numbers[1]
        }
    }

    return sum
}

func getGearNumbers(iGear int, row int) (bool, []int) {
    fmt.Printf("igear: %v, row: %v \n", iGear, row)
    var numbers []int
    var count int

    var line, prevLine, nextLine string
    var lineIdxs, prevIdxs, nextIdxs [][]int
    line = in[row]
    reNum := regexp.MustCompile("[0-9]+")
    lineIdxs = reNum.FindAllStringIndex(line, -1)
    var processNumber = func(digits []int, currentLine string) {
        count++
        num, _ := strconv.Atoi(currentLine[digits[0]:digits[1]])
        numbers = append(numbers, num)
    }
    for _, lineIdx := range lineIdxs {
        if lineIdx[1] == iGear {
            processNumber(lineIdx, line)
        }
        if iGear+1 == lineIdx[0] {
            processNumber(lineIdx, line)
        }
    }
    if row > 0 {
        prevLine = in[row-1]
        prevIdxs = reNum.FindAllStringIndex(prevLine, -1)
        for _, prevIdx := range prevIdxs {
            if prevIdx[1] == iGear {
                processNumber(prevIdx, prevLine)
            }
            if iGear+1 == prevIdx[0] {
                processNumber(prevIdx, prevLine)
            }
            if prevIdx[0] <= iGear && prevIdx[1] >= iGear+1 {
                processNumber(prevIdx, prevLine)
            }
        }
    }

    if row < len(in)-1 {
        nextLine = in[row+1]
        nextIdxs = reNum.FindAllStringIndex(nextLine, -1)
        for _, nextIdx := range nextIdxs {
            fmt.Printf("nextLine idxs: %v \n", nextIdx)
            if nextIdx[1] == iGear {
                processNumber(nextIdx, nextLine)
            }
            if iGear+1 == nextIdx[0] {
                processNumber(nextIdx, nextLine)
            }
            if nextIdx[0] <= iGear && nextIdx[1] >= iGear+1 {
                processNumber(nextIdx, nextLine)
            }
        }
    }
    if count != 2 {
        return false, numbers
    }

    return true, numbers
}

func countSum(line string, row int) int {
    var sum int
    reNum := regexp.MustCompile("[0-9]+")
    reSym := regexp.MustCompile("[^0-9a-zA-Z .]")
    numbers := reNum.FindAllStringIndex(line, -1)
    for _, number := range numbers {
        surrounding := getSurroundingString(number, row, line, in)
        isValidNumber := reSym.FindString(surrounding) != ""

        if isValidNumber {
            num  := line[number[0]:number[1]]
            numInt, _ := strconv.Atoi(num)
            sum += numInt
        }
    }

    return sum
}

func getSurroundingString(number []int, row int, line string, in []string) string {
    var str, top, bottom, topLeft, topRight, bottomLeft, bottomRight, left, right string
    var isFirstLine, isLastLine, isFirstChar, isLastChar bool
    var prevLine, nextLine string
    isFirstLine = row == 0
    isLastLine = row == len(in)-1
    isFirstChar = number[0] == 0
    isLastChar = number[1] == len(line)
    // fmt.Printf("check this out: %v, row: %v \n", number[1], line[number[0]:number[1]])
    if !isFirstLine {
        prevLine = in[row-1]
        top = prevLine[number[0]:number[1]]
        str += top
        if !isFirstChar {
            topLeft = string(prevLine[number[0]-1])
            str += topLeft
        }
        if !isLastChar {
            topRight = string(prevLine[number[1]])
            str += topRight
        }
    }
    if !isFirstChar {
        left = string(line[number[0]-1])
        str += left
    }
    if !isLastChar {
        right = string(line[number[1]])
        str += right
    }
    if !isLastLine {
        nextLine = in[row+1]
        bottom = nextLine[number[0]:number[1]]
        str += bottom
        if !isFirstChar {
            bottomLeft = string(nextLine[number[0]-1])
            str += bottomLeft
        }
        if !isLastChar {
            bottomRight = string(nextLine[number[1]])
            str += bottomRight
        }
    }

    return str
}

// search indexes of *
// loop indexes of *
// search and count for gears (bottom, left, top, right, topR, topL, bottomR, bottomL)
// if third gear found, break loop
// multiply the two numbers



