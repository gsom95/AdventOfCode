package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := flag.String("input", "input.txt", "path to the data")
	flag.Parse()

	input, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("cannot open %s: %v\n", *filename, err)
	}
	part1(input)
	input.Close()

	input, err = os.Open(*filename)
	if err != nil {
		log.Fatalf("cannot open %s: %v\n", *filename, err)
	}
	part2(input)
	input.Close()
}

func part1(input io.ReadCloser) {
	answer := 0

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	prevMeasurement, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("not a number read")
		return
	}
	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("not a number read")
			return
		}
		if measurement > prevMeasurement {
			answer += 1
		}
		prevMeasurement = measurement
	}

	fmt.Println(answer)
}

func part2(input io.ReadCloser) {
	answer := 0

	scanner := bufio.NewScanner(input)
	curSum := 0
	numbers := [3]int{}

	for i := 0; i < 3; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		numbers[i] = num
		curSum += num
	}

	indexToSub := 0
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		toExclude := numbers[indexToSub]
		if curSum < (curSum - toExclude + num) {
			answer += 1
		}
		numbers[indexToSub] = num
		indexToSub = (indexToSub + 1) % 3
	}

	fmt.Println(answer)
}
