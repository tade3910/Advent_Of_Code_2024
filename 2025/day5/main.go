package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type idRange struct {
	Max int
	Min int
}

func getInt(v string) int {
	if v == "" {
		return 0
	}
	val, err := strconv.Atoi(v)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func part1(file_name string) {
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validRanges := []*idRange{}
	isRange := true
	numValids := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isRange = false
			continue
		}
		line = strings.TrimSpace(line)
		if isRange {
			splitLine := strings.Split(line, "-")
			start, end := getInt(splitLine[0]), getInt(splitLine[1])
			validRanges = append(validRanges, &idRange{
				Max: end,
				Min: start,
			})
		} else {
			id := getInt(line)
			for _, validRange := range validRanges {
				if id <= validRange.Max && id >= validRange.Min {
					numValids++
					break
				}
			}
		}
	}
	fmt.Printf("There are: %d valid ids\n", numValids)
}

func getValidRanges(file_name string) []*idRange {
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validRanges := []*idRange{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			return validRanges
		}
		line = strings.TrimSpace(line)
		splitLine := strings.Split(line, "-")
		start, end := getInt(splitLine[0]), getInt(splitLine[1])
		validRanges = append(validRanges, &idRange{
			Max: end,
			Min: start,
		})
	}
	return validRanges
}

func shrinkRanges(validRanges []*idRange) []*idRange {
	shrunk_ranges := []*idRange{}
	for _, curRuange := range validRanges {
		start, end := curRuange.Min, curRuange.Max
		for _, validRange := range validRanges {
			if curRuange == validRange {
				//skip same pointer
				continue
			}
			if start < validRange.Min && end >= validRange.Min {
				//Extend left
				validRange.Min = start
			}
			if end > validRange.Max && start <= validRange.Max {
				//Extend Right
				validRange.Max = end
			}
		}
	}
	//We write over ranges all the time, use map to remove duplicates
	range_map := make(map[idRange]bool)
	for _, validRange := range validRanges {
		_, ok := range_map[*validRange]
		if !ok {
			range_map[*validRange] = true
			shrunk_ranges = append(shrunk_ranges, validRange)
		}
	}
	return shrunk_ranges
}

func part2(file_name string) {
	numValids := 0
	for _, curRuange := range shrinkRanges(getValidRanges(file_name)) {
		numValids += (curRuange.Max - curRuange.Min + 1)
	}
	fmt.Printf("There are: %d valid ids\n", numValids)
}

func main() {
	part1("sample.txt")
	part1("input.txt")
	part2("sample.txt")
	part2("input.txt")
}

/*
//After coding I realized I am stupid
//My solution is O(n^2) could be O(nlgn)

Step 1:
Get all ranges -> O(n)
Sort ranges by start time -> O(nlgn)
Iterate through array growing ranges until you get to a range that can't be joined in current range
If we're at a spot that the min of cur > max grown array i.e
1-5 -> 1-10
2-10 -> Not added
11- 25 -> 11-25
We only need to pass through once and only insert when at the end or next element is bigger
*/
