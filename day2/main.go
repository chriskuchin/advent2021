package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var x int
	var y int
	var aim int
	for scanner.Scan() {
		currentVal := scanner.Text()
		resultX, resultY, resultAim := processInstruction(currentVal, aim)
		x += resultX
		y += resultY
		aim += resultAim
	}

	fmt.Println(x, y)
	fmt.Println(x * y)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

// down X increases your aim by X units.
// up X decreases your aim by X units.
// forward X does two things:
// 	It increases your horizontal position by X units.
// 	It increases your depth by your aim multiplied by X.
func processInstruction(rawInstruction string, currentAim int) (int, int, int) {
	var x int
	var y int
	var aim int
	instruction := strings.Split(rawInstruction, " ")
	magnitude, _ := strconv.Atoi(instruction[1])

	if instruction[0] == "up" {
		aim = -1 * magnitude
	} else if instruction[0] == "down" {
		aim = 1 * magnitude
	} else if instruction[0] == "forward" {
		x = 1 * magnitude
		y = 1 * magnitude * currentAim
	} else if instruction[0] == "back" {
		x = -1 * magnitude
		y = -1 * magnitude * currentAim
	}

	return x, y, aim
}
