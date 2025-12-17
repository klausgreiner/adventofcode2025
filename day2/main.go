package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split line by commas and add each part as a separate element
		line := scanner.Text()
		parts := make([]string, 0)
		start := 0
		for i, c := range line {
			if c == ',' {
				parts = append(parts, line[start:i])
				start = i + 1
			}
		}
		// Add the last part if the line does not end with ','
		if start < len(line) {
			parts = append(parts, line[start:])
		}
		// Filter out empty strings (for trailing commas)
		for _, p := range parts {
			if p != "" {
				lines = append(lines, p)
			}
		}
	}
	return lines
}

func main() {
	var sum int = 0
	lines := readInput()
	for _, line := range lines {
		// split line by '-' transform to ints
		// create a list with the range of the two ints
		// sort the list
		// verify which values from list are invalid
		// being invalid means being equal for example 11-22 2 == 2 1 == 1
		// 99 is invalid
		//1188511885 is invalida cuz 11885 == 11885
		values := strings.Split(line, "-")
		minInt, _ := strconv.Atoi(values[0])
		maxInt, _ := strconv.Atoi(values[1])

		for maxInt >= minInt {
			//fmt.Println("current value", maxInt)
			// convert maxInt to a list of its digits
			digits := []int{}
			for _, ch := range strconv.Itoa(maxInt) {
				digits = append(digits, int(ch-'0'))
			}
			// Check if the digits form two identical consecutive sequences (e.g., 2323, 11881188, etc.)
			var areEqual bool = false
			if len(digits)%2 == 0 {
				half := len(digits) / 2
				areEqual = true
				for i := 0; i < half; i++ {
					if digits[i] != digits[i+half] {
						areEqual = false
						break
					}
				}

				if areEqual {
					sum += maxInt
					//fmt.Println("last accmulation was", maxInt)
				}

			}

			maxInt -= 1

		}
		//fmt.Println("curren sum is ", sum)

	}
	fmt.Println("result is", sum)

}
