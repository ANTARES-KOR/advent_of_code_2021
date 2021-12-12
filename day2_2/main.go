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
	aim := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		words := strings.Split(line, " ")

		inst := words[0]
		number, _ := strconv.Atoi(words[1])

		fmt.Println(depth, horizontal, aim)

		switch inst {
		case "forward":
			horizontal += number
			depth += number * aim
		case "down":
			aim += number
		case "up":
			aim -= number
		}
	}
	fmt.Println(depth * horizontal)
}
