package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []int
	for scanner.Scan() {
		currentVal, _ := strconv.Atoi(scanner.Text())

		input = append(input, currentVal)
	}

	fmt.Println(processIncreasing(calculateRollingWindow(input)))

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}

func calculateRollingWindow(input []int) []int {
	var groups []int
	for key, val := range input {
		if key+1 < len(input) && key+2 < len(input) {
			groups = append(groups, val+input[key+1]+input[key+2])
		}
	}

	return groups
}

func processIncreasing(input []int) int {
	var previousVal int
	var count int
	for _, val := range input {
		currentVal := val
		if previousVal != 0 {
			if currentVal > previousVal {
				count++
			}
		}

		previousVal = currentVal
	}

	return count
}
