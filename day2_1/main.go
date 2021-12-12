package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)

	depth := 0
	horizontal := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		words := strings.Split(line, " ")

		inst := words[0]
		number, _ := strconv.Atoi(words[1])

		switch inst {
		case "forward":
			horizontal += number
		case "down":
			depth += number
		case "up":
			depth -= number
		}
	}
	fmt.Println(depth * horizontal)
}
