package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// create map reference
// split input lines
// split color balls
// get number and color
// match color number with map reference


var bagOfCubes = map[string]int{
    "red": 12,
    "green": 13,
    "blue": 14,
}

func main() {
    input, _ := os.ReadFile("day2_in.txt")
    games := strings.Split(string(input), "\n")
    Sum(games[:100])
    SumV2(games[:100])
}

func Sum(games []string) {
    var sum int
    for i, game := range games {
        isPossibleGame := true
        cubeString := strings.Split(game, ": ")[1]
        cubeSets := strings.Split(cubeString, "; ")
        for _, cubeSet := range cubeSets {
            if !isPossibleSet(cubeSet) {
                isPossibleGame = false
                break
            }
        }

        if isPossibleGame {
            sum += i + 1
        }


        // fmt.Printf("Current Sum: %v \n", sum)
    }

    fmt.Printf("Total Sum: %v \n", sum)
}

func SumV2(games []string) {
    var sum int
    for _, game := range games {
        cubeString := strings.Split(game, ": ")[1]
        sum += CheckMinimumSet(cubeString)
    }
    fmt.Printf("SumV2: %v \n", sum)
}

func CheckMinimumSet(cubeString string) int {
    var minimum = map[string]int{
        "red": 0,
        "green": 0,
        "blue": 0,
    }

    cubeSets := strings.Split(cubeString, "; ")
    for _, cubeSet := range cubeSets {
        cubes := strings.Split(cubeSet, ", ")
        for _, cube := range cubes {
            numColorPair := strings.Split(cube, " ")
            num, _ := strconv.Atoi(numColorPair[0])
            color := numColorPair[1]
            if minimum[color] < num {
                minimum[color] = num
            }
        }
    }

    return minimum["red"]*minimum["green"]*minimum["blue"]
}

func isPossibleSet(set string) bool {
    cubes := strings.Split(set, ", ")
    for _, cube := range cubes {
        numColorPair := strings.Split(cube, " ")
        if num, _ := strconv.Atoi(numColorPair[0]); num > bagOfCubes[numColorPair[1]] {
            return false
        }
    }

    return true
}
