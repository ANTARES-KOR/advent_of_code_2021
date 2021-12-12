package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("./input.txt")
	if err != nil  {
		panic(err)
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)

	increaseCnt := 0;
	cnt := 0;
	var number int
	var mem = [3]int{0, 0, 0}

	for fileScanner.Scan() {
		number, _ = strconv.Atoi(fileScanner.Text())

		if cnt >= 3 {
			if number > mem[0] {
				increaseCnt++
			}
		}

		mem[0] = mem[1]
		mem[1] = mem[2]
		mem[2] = number
		cnt++
	}

	fmt.Println(increaseCnt)

}