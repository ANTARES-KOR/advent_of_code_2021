package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	input, err := os.Open("./input.txt")
	// go에서 특정 자료형의 null 을 나타내는 값
	if err != nil {
		panic(err)
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)

	var prev = 1000000;
	var new int;
	cnt := 0

	for fileScanner.Scan() {
		new, _ = strconv.Atoi(fileScanner.Text())
		if new > prev {
			cnt++;
		}
		prev = new;
	}

	fmt.Println(cnt)
	
}