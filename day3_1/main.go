package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	var zeroCount = [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var oneCount = [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	reportCount := 0

	fileScanner := bufio.NewScanner(input)
	for fileScanner.Scan() {
		reportCount++
		new := fileScanner.Text()

		for i, rune := range new {
			if rune == '0' {
				zeroCount[i]++
			} else {
				oneCount[i]++
			}
		}
	}

	var gamma = ""
	var epsilon = ""

	for i := 0; i < 12; i++ {
		if zeroCount[i] > oneCount[i] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	fmt.Println(gamma, epsilon)

	gammaNum, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonNum, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(gammaNum * epsilonNum)

}
