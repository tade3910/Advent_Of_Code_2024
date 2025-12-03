package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func string_to_int(v string) int {
	if v == "" {
		return 0
	}
	val, err := strconv.Atoi(v)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func largest_voltage(bank string) int { //short cut for part 1
	first, second := 0, 0
	for index, r := range bank {
		cur := int(r - '0')
		if cur > first && index < len(bank)-1 {
			first = cur
			second = 0
			//Unless can't be in 10 place so try 1
		} else if cur > second {
			second = cur // replace 1 place
		}
	}
	voltage_str := strconv.Itoa(first) + strconv.Itoa(second)
	voltage := string_to_int(voltage_str)
	// fmt.Printf("Largest is: %d, Followed by: %d with voltage: %d\n", first, second, voltage)
	return voltage
}

func find_largest_range(bank []int, start_index, place int) (int, int) {
	//place = 2 10s
	//place =3 100s
	max_index := len(bank) - place
	cur_index, place_index, largest := start_index+1, start_index, bank[start_index]
	for cur_index <= max_index {
		cur := bank[cur_index]
		if cur > largest {
			place_index = cur_index
			largest = cur
		}
		cur_index++
	}
	return largest, place_index
}

func get_voltage(bank_string string, batteries int) int {
	bank := make([]int, len(bank_string))
	for index, r := range bank_string {
		cur := int(r - '0')
		bank[index] = cur
	}
	voltage := 0
	start_index := 0
	for place := batteries; place > 0; place-- {
		largest, place_index := find_largest_range(bank, start_index, place)
		start_index = place_index + 1
		voltage += largest * int(math.Pow10(place-1))
	}
	return voltage
}

func runner(file_name string, part2 bool) int {
	batteries := 2
	if part2 {
		batteries = 12
	}
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total_voltage := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		total_voltage += get_voltage(line, batteries)
	}
	fmt.Printf("Total voltage is %d\n", total_voltage)
	return total_voltage
}

func main() {
	runner("sample.txt", false)
	runner("input.txt", false)
	runner("sample.txt", true)
	runner("input.txt", true)
	fmt.Printf("%d\n", largest_voltage("7176")) // Edge case didn't initally think of

}
