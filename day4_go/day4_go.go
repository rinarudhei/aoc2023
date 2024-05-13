package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)


func main() {
    inBytes, _ := os.ReadFile("day4_in.txt")
    cardsBytes := bytes.Split(inBytes, []byte("\n"))
    var sum int
    for _, cardBytes := range cardsBytes {
        if len(cardBytes) == 0 {
            continue
        }
        lineBytes := strings.Split(string(cardBytes), ": ")[1]
        sum += CountPoint(string(lineBytes))
    }

    fmt.Printf("result: %v \n", sum)
}

func CountPoint(lineString string) int {
    fmt.Printf("linestring: %s \n", lineString)
    re := regexp.MustCompile("[0-9]+")
    numbersString := re.FindAllString(lineString, -1)
    winningNums := numbersString[:10]
    numsWeHave := numbersString[10:]
    var counter int
    for _, winningNum := range winningNums {
        for _, numWeHave := range numsWeHave {
            if winningNum == numWeHave {
                if counter == 0 {
                    counter += 1
                } else {
                    counter *= 2
                }
                break
            }
        }
    }
    fmt.Printf("counter: %v \n", counter)
    return counter
}

//split bytes by \n
//loop it, get all the numbers in array
//compare every 1~10th numbers with 11th numbers and above
//count the point


