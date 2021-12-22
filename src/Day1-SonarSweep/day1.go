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

	result1 := part1(array)
	fmt.Println(fmt.Sprintf("part 1 result: %d", result1))
	result2 := part2(array)
	fmt.Println(fmt.Sprintf("part 2 result: %d", result2))
}

func part1(array []string) int {
	part1Result := 0
	previousValue, _ := strconv.Atoi(array[0]);
	//fmt.Println(previousValue)
	//fmt.Println(err)

	for _, s := range array {
		currentNumber, _ := strconv.Atoi(s)
		// Should handle the error here.
		if currentNumber > previousValue {
			part1Result++
			//fmt.Println(s + " Higher than " + previousValue)
		}
		previousValue = currentNumber
	}

	return part1Result
}

func part2(array []string) int {
	part2Result := 0

	v1, _ := strconv.Atoi(array[0])
	v2, _ := strconv.Atoi(array[1])
	v3, _ := strconv.Atoi(array[2])
	previousValue := v1 + v2 + v3
	for i := 1; i < len(array) - 2; i++ {
		// Sum the first value
		v1, _ = strconv.Atoi(array[i])
		v2, _ = strconv.Atoi(array[i+1])
		v3, _ = strconv.Atoi(array[i+2])
		currentNumber := v1 + v2 + v3
		if currentNumber > previousValue {
			part2Result++
			//fmt.Println(s + " Higher than " + previousValue)
		}
		previousValue = currentNumber
	}

	return part2Result
}