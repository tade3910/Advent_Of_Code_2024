package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func get_int(v string) int {
	if v == "" {
		return 0
	}
	val, err := strconv.Atoi(v)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func invalid_ids(start_string, stop_string string, repeats int) []int {
	/**
	Idea is to check
	x x
	xx xx
	xxx xxx

	If start is
	abc -> We can't do anything with this so add 1 to our mid
	ab ab -> check if withing stop

	min half for split_len = 1 -> 1 == 10 ^ 0 max = 99 -> 10^1 - 1
	min half for split_len = 2 -> 10 == 10 ^ 1 max = 999 -> 10 ^2 -1

	Check if cur greater than min and end when greater than stop
	*/
	invalid_slice := []int{}
	start := get_int(start_string)
	stop := get_int(stop_string)

	for length := len(start_string); length <= len(stop_string); length++ {
		if length%repeats != 0 {
			continue // invalid IDs must have divisible length
		}

		split_len := length / repeats
		min_half := int(math.Pow10(split_len - 1))
		max_half := int(math.Pow10(split_len)) - 1

		for half := min_half; half <= max_half; half++ {
			halfStr := strconv.Itoa(half)
			curStr := ""
			for range repeats {
				curStr += halfStr
			}
			cur := get_int(curStr)

			if cur > stop {
				break
			}
			if cur >= start {
				invalid_slice = append(invalid_slice, cur)
			}
		}
	}

	return invalid_slice
}

func variable_invalid_ids(start_string, stop_string string) int {
	sums := make(map[int]bool)
	for divisor := 2; divisor <= len(stop_string); divisor++ {
		sums_slice := invalid_ids(start_string, stop_string, divisor)
		//Return slice cause we need to know distinct items returned
		//Double counted values could have same result
		for _, val := range sums_slice {
			sums[val] = true
		}
	}
	sum := 0
	for cur := range sums {
		if cur != 0 {
			sum += cur
		}
	}
	return sum
}

func sum_slice(slice []int) int {
	sum := 0
	for _, cur := range slice {
		sum += cur
	}
	return sum
}

func part_1(file_name string) {
	data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), ",")
	solution := 0
	for _, id_ranges := range input {
		split_range := strings.Split(strings.TrimSpace(id_ranges), "-")
		solution += sum_slice(invalid_ids(split_range[0], split_range[1], 2))
	}
	fmt.Printf("Solution is %d\n", solution)

}

func part_2(file_name string) {
	data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), ",")
	solution := 0
	for _, id_ranges := range input {
		split_range := strings.Split(strings.TrimSpace(id_ranges), "-")
		solution += variable_invalid_ids(split_range[0], split_range[1])
	}
	fmt.Printf("Solution is %d\n", solution)

}
func main() {
	part_1("sample.txt")
	part_1("input.txt")
	part_2("sample.txt")
	part_2("input.txt")

}
