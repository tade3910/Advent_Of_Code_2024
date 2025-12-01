package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func get_password(file_name string) int {
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	negative_byte := byte('L')
	cur, password := 50, 0
	for scanner.Scan() {
		line := scanner.Text()
		sign := line[0]
		multiplier := 1
		if negative_byte == sign {
			multiplier = -1
		}
		magnitude, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		cur = get_val(cur + (multiplier * magnitude))
		// fmt.Printf("Cur is %d\n", cur)
		if cur == 0 {
			password++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return password
}
func get_val(val int) int {
	/**
	val % 100 -> If positive cool
	if negative need to add len i.e -5 + 100 = 95
	but will be 5 + 100 = 105
	just modulo to normalize
	95 % 100 = 95
	105 % 100 = 5
	*/
	return (val%100 + 100) % 100
}

func get_fancy_password(file_name string) int {
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	negative_byte := byte('L')
	cur, password := 50, 0

	for scanner.Scan() {
		line := scanner.Text()
		sign := line[0]
		multiplier := 1
		if negative_byte == sign {
			multiplier = -1
		}

		magnitude, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		// Count how many times we pass 0
		for range magnitude { // Was trying to be too smart just brute force
			cur += multiplier
			//wrap around logic
			switch cur {
			case -1:
				cur = 99
			case 100:
				cur = 0
			}
			if cur == 0 {
				password++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return password
}

func main() {
	password := get_password("example.txt")
	fmt.Printf("password is %d\n", password)
	password = get_password("first.txt")
	fmt.Printf("password is %d\n", password)
	password = get_fancy_password("example.txt")
	fmt.Printf("password is %d\n", password)
	password = get_fancy_password("first.txt")
	fmt.Printf("password is %d\n", password)

}
