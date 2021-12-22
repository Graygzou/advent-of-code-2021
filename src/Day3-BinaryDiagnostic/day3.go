package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {

	dat, _ := os.ReadFile("./example.txt")

	//fmt.Print(string(dat))
	//fmt.Println(string(dat))
	
	strInput := string(dat)
	array := strings.Split(strInput, "\n")
	fmt.Println(array)

	result := part1(array)
	fmt.Println(fmt.Sprintf("part 1 result: %d", result))

	result = part2(array)
	fmt.Println(fmt.Sprintf("part 2 result: %d", result))
}

func part1(array []string) (int64) {
	arrayLen := len(array)

	list := make([]int, len(array[0]))
	for _, s := range array {
		for i := 0; i < len(s); i++ {
			intVar, _ := strconv.Atoi(s[i:i+1])
			list[i] = list[i] + intVar
		}
	}

	mostCommonBits := ""
	leastCommonBits := ""
	for _, v := range list {
		if v >= arrayLen/2 {
			mostCommonBits = mostCommonBits + "1"
			leastCommonBits = leastCommonBits + "0"
		} else {
			mostCommonBits = mostCommonBits + "0"
			leastCommonBits = leastCommonBits + "1"
		}
	}

	gammaRate, _ := strconv.ParseInt(mostCommonBits, 2, 64)
	fmt.Println(gammaRate)

	powerConsumption, _ := strconv.ParseInt(leastCommonBits, 2, 64)
	fmt.Println(powerConsumption)

	return gammaRate * powerConsumption
}

func part2(array []string) (int64) {
	list := make([]string, len(array[0]))
	copy(list, array)
	
	bitIndex := 0
	for len(list) > 0 && bitIndex < len(array[0]) {
		mostCommonBit := FindMostCommonBitAtIndex(array, bitIndex)
		fmt.Println("Intermediate value")
		fmt.Println(mostCommonBit)

		for i := 0; i < len(list); i++ {
			// Remove it from the list if the bit does not contains the most common bit
			currentBit, _ := strconv.Atoi(list[i][bitIndex:bitIndex+1])
			if currentBit != mostCommonBit {
				list = append(list[:i], list[i+1:]...)
				i--
			}
		}

		fmt.Println("Intermediate list")
		fmt.Println(list)

		bitIndex++
	}
	
	mostCommonBits := ""
	leastCommonBits := ""

	gammaRate, _ := strconv.ParseInt(mostCommonBits, 2, 64)
	fmt.Println(gammaRate)

	powerConsumption, _ := strconv.ParseInt(leastCommonBits, 2, 64)
	fmt.Println(powerConsumption)

	return gammaRate * powerConsumption
}

func FindMostCommonBitAtIndex(array []string, index int) (int) {
	bitSum := 0

	for _, s := range array {
		intVar, _ := strconv.Atoi(s[index:index+1])
		bitSum = bitSum + intVar
	}

	result := 0
	if bitSum > len(array)/2 {
		result = 1
	}

	return result
}