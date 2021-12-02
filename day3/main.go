package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentVal := scanner.Text()
		fmt.Println(currentVal)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
