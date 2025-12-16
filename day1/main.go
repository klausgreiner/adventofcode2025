package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		lines = append(lines, scanner.Text())
	}
	return lines
}

// math euclidean modulo
func PosMod(n, m int) int {
	return ((n % m) + m) % m
}

func main() {
	var currentPosition int = 50
	var secretPassword int = 0
	lines := readInput()
	for _, line := range lines {
		fmt.Println(line)
		direction := line[0]
		number, _ := strconv.Atoi(line[1:])
		var result int
		if direction == 'L' {
			result = PosMod(currentPosition-number, 100)
		} else if direction == 'R' {
			result = PosMod(currentPosition+number, 100)
		}
		currentPosition = result
		fmt.Printf("currentPosition: %d\n", currentPosition)
		if currentPosition == 0 {
			secretPassword += 1
			fmt.Println("Secret password incremented to:", secretPassword)
		}
	}
	fmt.Println("Secret password:", secretPassword)

}
