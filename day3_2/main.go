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

	var zeroCount = [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var oneCount = [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	reportCount := 0

	var reports = make([]string, 10000)

	fileScanner := bufio.NewScanner(input)
	for fileScanner.Scan() {
		new := fileScanner.Text()
		reports[reportCount] = new
		reportCount++
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

	var oxygenPrefixes [12] string
	var carbonPrefixes [12] string
	var oxygenMatchCount[12] int
	var carbonMatchCount[12] int
	var oxygenRatingCandidates[12] string
	var carbonRatingCandidates[12] string

	for i:=0; i<12; i++ {
		oxygenPrefixes[i] = gamma[:12-i]
		carbonPrefixes[i] = epsilon[:12-i]
		oxygenMatchCount[i] = 0
		carbonMatchCount[i] = 0
		oxygenRatingCandidates[i] = ""
		carbonRatingCandidates[i] = ""
	}


	for i:=0; i<reportCount; i++ {
		for j:=0; j<12; j++ {
			if strings.HasPrefix(reports[i], oxygenPrefixes[j]) {
				oxygenMatchCount[j]++
				oxygenRatingCandidates[j] = reports[i]
			}
			if strings.HasPrefix(reports[i], carbonPrefixes[j]) {
				carbonMatchCount[j]++
				carbonRatingCandidates[j] = reports[i]
			}
		}
	}

	var oxygenRating, carbonRating string

	for i:=0; i<12; i++ {
		if oxygenMatchCount[i] == 1 {
			oxygenRating = oxygenRatingCandidates[i]
		}
		if carbonMatchCount[i] == 1 {
			carbonRating = carbonRatingCandidates[i]
		}
	}

	fmt.Println(oxygenRating, carbonRating)

	oxygenRatingNum, _ := strconv.ParseInt(oxygenRating, 2, 64)
	carbonRatingNum, _ := strconv.ParseInt(carbonRating, 2, 64)

	fmt.Println(oxygenRatingNum * carbonRatingNum)

}
