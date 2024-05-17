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
        res, _ := CountPoint(string(lineBytes))
        sum += res
    }
    fmt.Printf("result: %v \n", sum)

    var cardsDeck [212]int
    var countCardsDeck [212]int
    for i, _ := range cardsDeck[:len(cardsDeck)] {
        cardsDeck[i] = -1
    }
    CountStretchCards(cardsBytes, &cardsDeck, 0,&countCardsDeck, false)
    var sumCard int
    for i, _ := range countCardsDeck {
        sumCard += countCardsDeck[i]
    }
    fmt.Printf("result: %v \n", sumCard)
}

func CountStretchCards(cardsBytes [][]byte, cardsDeck *[212]int, idx int, countCardsDeck *[212]int, isBonus bool) {
    if idx >= len(cardsBytes)-1 {
        return
    }
    countCardsDeck[idx]++
    lineBytes := strings.Split(string(cardsBytes[idx]), ": ")[1]
    var bonusCount int
    if cardsDeck[idx] == -1 {
        _, increment := CountPoint(string(lineBytes))
        bonusCount = increment
        cardsDeck[idx] = increment
    } else {
        bonusCount = cardsDeck[idx]
    }

    if bonusCount > 0 {
        for j := 1; j <= bonusCount; j++ {
            CountStretchCards(cardsBytes, cardsDeck, idx+j, countCardsDeck, true)
        }
    }

    if !isBonus {
        CountStretchCards(cardsBytes, cardsDeck, idx+1, countCardsDeck, false)
    }
}

func CountPoint(lineString string) (int, int) {
    re := regexp.MustCompile("[0-9]+")
    numbersString := re.FindAllString(lineString, -1)
    winningNums := numbersString[:10]
    numsWeHave := numbersString[10:]
    increment := 0
    var counter int
    for _, winningNum := range winningNums {
        for _, numWeHave := range numsWeHave {
            if winningNum == numWeHave {
                increment++
                if counter == 0 {
                    counter += 1
                } else {
                    counter *= 2
                }
                break
            }
        }
    }
    return counter, increment
}

//split bytes by \n
//loop it, get all the numbers in array
//compare every 1~10th numbers with 11th numbers and above
//count the point


