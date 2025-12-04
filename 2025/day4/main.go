package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func count_left_right(count_floor [][]int, row, col int) {
	//left
	if col > 0 {
		count_floor[row][col-1]++
	}
	//right
	if col < len(count_floor[row])-1 {
		count_floor[row][col+1]++
	}
}

func count_adjacents(count_floor [][]int, row, col int) {
	top_row, bottom_row := row-1, row+1
	//Update up
	if top_row >= 0 {
		count_floor[top_row][col]++
		count_left_right(count_floor, top_row, col)
	}
	//update lr
	count_left_right(count_floor, row, col)
	//update bottom
	if bottom_row < len(count_floor) {
		count_floor[bottom_row][col]++
		count_left_right(count_floor, bottom_row, col)
	}
}

func num_accesible(floor_map []string) (int, []string) {
	const tp = '@'
	const none = '.'
	const hit = 'x'
	//Initialize count floor
	count_floor := make([][]int, len(floor_map))
	visual_map := make([][]rune, len(floor_map))
	for i, row := range floor_map {
		count_floor[i] = make([]int, len(row))
		visual_map[i] = make([]rune, len(row))
	}
	//Visualizing loop, could prolly do all in one
	//Thank God I did it like this at first, don't need to rewrite for part 2
	for row_index, row := range floor_map {
		for col_index, col := range row {
			if col != tp {
				continue
			}
			count_adjacents(count_floor, row_index, col_index)
		}
	}
	accessible := 0
	for row_index, row := range floor_map {
		for col_index, col := range row {
			if col == tp {
				if count_floor[row_index][col_index] < 4 {
					accessible++
					visual_map[row_index][col_index] = hit
				} else {
					visual_map[row_index][col_index] = tp
				}
			} else {
				visual_map[row_index][col_index] = none
			}
		}
	}
	next_world := make([]string, len(floor_map))
	for index, row := range visual_map {
		next_world[index] = string(row)
	}
	return accessible, next_world
}

func part_1(file_name string) {
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	floor_map := []string{}
	for scanner.Scan() {
		floor_map = append(floor_map, strings.TrimSpace(scanner.Text()))
	}
	accessible, _ := num_accesible(floor_map)
	fmt.Printf("Num accessible is %d\n", accessible)
}

func part_2(file_name string) {
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	floor_map := []string{}
	for scanner.Scan() {
		floor_map = append(floor_map, strings.TrimSpace(scanner.Text()))
	}
	total_accssible := 0
	accessible, next_world := num_accesible(floor_map)
	for accessible != 0 {
		total_accssible += accessible
		// fmt.Printf("Num accessible is %d\n", accessible)
		accessible, next_world = num_accesible(next_world)
	}
	fmt.Printf("Total accessible is %d\n", total_accssible)

}

func main() {
	part_1("sample.txt")
	part_1("input.txt")
	part_2("sample.txt")
	part_2("input.txt")
}
