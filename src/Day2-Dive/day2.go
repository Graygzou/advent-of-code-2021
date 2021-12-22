package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {

	dat, _ := os.ReadFile("./input.txt")

	//fmt.Print(string(dat))
	//fmt.Println(string(dat))
	
	strInput := string(dat)
	array := strings.Split(strInput, "\n")
	fmt.Println(array)

	horizontal, depth := part1(array)
	fmt.Println(fmt.Sprintf("part 1 result: %d", horizontal * depth))

	horizontal, depth = part2(array)
	fmt.Println(fmt.Sprintf("part 2 result: %d", horizontal * depth))
}

func part1(array []string) (int, int) {
	horizontal := 0
	depth := 0

	for _, s := range array {
		currentInstruction := strings.Split(s, " ");
		if len(currentInstruction) > 1 {
			value, _ := strconv.Atoi(currentInstruction[1]);
			//fmt.Println(err)

			fmt.Println(fmt.Sprintf("%s and %d ", currentInstruction[0][0:1], value))

			switch currentInstruction[0][0:1] {
				case "f", "F": // Forward
					horizontal += value
				case "d", "D": // Down
					depth += value
				case "u", "U": // Up
					depth -= value
			}
		} else {
			fmt.Println(fmt.Sprintf("instruction unknow: %s", currentInstruction))
		}
	}

	return horizontal, depth
}

func part2(array []string) (int, int) {
	horizontal := 0
	depth := 0
	aim := 0

	for _, s := range array {
		currentInstruction := strings.Split(s, " ");
		if len(currentInstruction) > 1 {
			value, _ := strconv.Atoi(currentInstruction[1]);
			//fmt.Println(err)

			fmt.Println(fmt.Sprintf("%s and %d ", currentInstruction[0][0:1], value))

			switch currentInstruction[0][0:1] {
				case "f", "F": // Forward
					horizontal += value
					depth += aim * value
				case "d", "D": // Down
					aim += value;
				case "u", "U": // Up
					aim -= value;
			}
		} else {
			fmt.Println(fmt.Sprintf("instruction unknow: %s", currentInstruction))
		}
	}

	return horizontal, depth
}