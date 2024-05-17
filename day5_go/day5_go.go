package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const INPUT_FILE = "day5_in.txt"
const SEED_TO_SOIL_MAP = "seed-to-soil map"
var res, _ = os.ReadFile(INPUT_FILE)
var lines = bytes.Split(res, []byte("\n"))
var reNum = regexp.MustCompile("[0-9]+")
var re = regexp.MustCompile("(\n[a-z :-]+(\n))")

func main() {

    seeds, maps := GetMaps()
    fmt.Printf("check seedsd: %v \n", seeds)
    fmt.Printf("check maps: %v \n", maps)

    minLoc := GetMinimumLocation(seeds, maps)
    fmt.Printf("check min location: %v \n", minLoc)
}

func GetMinimumLocation(seeds []int, maps [][][]int) (minLoc int) {
    for _, s := range seeds {
        if loc := CheckDestination(s, maps); minLoc > loc || minLoc == 0 {
            minLoc = loc
        }
    }

    return
}

func CheckDestination(seed int, maps [][][]int) int {
    for _, m := range maps {
        for _, p := range m {
            destination := p[0]
            source := p[1]
            r := p[2]

            if source+r >= seed && seed >= source {
                diff := destination-source
                seed += diff
                break
            }

        }
    }

    return seed
}

func GetSeeds(firstLine string) []int {
    seedsString := reNum.FindAllString(firstLine, -1)
    var seeds []int
    for _, seedString := range seedsString {
        seed, _ := strconv.Atoi(seedString)
        seeds = append(seeds, seed)
    }

    return seeds
}

func GetMaps() (seeds []int, maps [][][]int){
    mapsString := re.Split(string(res), -1)
    seeds = GetSeeds(mapsString[0])
    for _, mapString := range mapsString[1:] {
        var m [][]int
        s := strings.Split(mapString, "\n")
        for _, c := range s {
            if len(c) > 0 {
                numsString := reNum.FindAllString(c, -1)
                var nums []int
                for _, n := range numsString {
                    res, _ := strconv.Atoi(n)
                    nums = append(nums, res)
                }
                m = append(m, nums)
            }
        }

        maps = append(maps, m)
    }

    return
}
