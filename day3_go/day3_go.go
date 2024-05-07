package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func main() {
    ins, _ := os.ReadFile("day3_in.txt")
    in := strings.Split(string(ins), "\n")
    in = in[0:140]
    reNum := regexp.MustCompile("[0-9]+")
     reSym := regexp.MustCompile("[^0-9a-zA-Z .]")
    var sum int
    for row, line := range in {
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
    }

    fmt.Printf("check sum: %v \n", sum)
}

func getSurroundingString(number []int, row int, line string, in []string) string {
    var str, top, bottom, topLeft, topRight, bottomLeft, bottomRight, left, right string
    var isFirstLine, isLastLine, isFirstChar, isLastChar bool
    var prevLine, nextLine string
    isFirstLine = row == 0
    isLastLine = row == len(in)-1
    isFirstChar = number[0] == 0
    isLastChar = number[1] == len(line)
    fmt.Printf("check this out: %v, row: %v \n", number[1], line[number[0]:number[1]])
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



