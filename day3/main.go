package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := flag.String("input", "input.txt", "path to the data")
	flag.Parse()

	input, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("cannot open %s: %v\n", *filename, err)
	}
	part1(input)
	if err := input.Close(); err != nil {
		log.Println(err)
	}

	input, err = os.Open(*filename)
	if err != nil {
		log.Fatalf("cannot open %s: %v\n", *filename, err)
	}
	part2(input)
	if err := input.Close(); err != nil {
		log.Println(err)
	}
}

func part1(input io.ReadCloser) {
	scanner := bufio.NewScanner(input)
	var amntOfOnes []int

	lines := 0
	for scanner.Scan() {
		lines++
		text := scanner.Text()
		if amntOfOnes == nil {
			amntOfOnes = make([]int, len(text))
		}

		for i, digit := range text {
			if digit == '1' {
				amntOfOnes[i]++
			}
		}
	}

	threshold := lines / 2
	var gammaStr strings.Builder
	var epsilonStr strings.Builder
	for _, a := range amntOfOnes {
		if a >= threshold {
			gammaStr.WriteString("1")
			epsilonStr.WriteString("0")
		} else {
			gammaStr.WriteString("0")
			epsilonStr.WriteString("1")
		}
	}
	gamma, _ := strconv.ParseInt(gammaStr.String(), 2, 0)
	epsilon, _ := strconv.ParseInt(epsilonStr.String(), 2, 0)

	fmt.Println(gamma * epsilon)
}

func part2(input io.ReadCloser) {
	scanner := bufio.NewScanner(input)

	var numbers []string
	lines := 0
	for scanner.Scan() {
		lines++
		text := scanner.Text()
		numbers = append(numbers, text)
	}

	amountOfBits := len(numbers)
	curNumbers := numbers
	for curBit := 0; curBit < amountOfBits; curBit++ {
		if len(curNumbers) == 1 {
			break
		}
		amountOfOnes := 0

		onesNumbers := []string{}
		zeroNumbers := []string{}
		for _, number := range curNumbers {
			if number[curBit] == '1' {
				amountOfOnes++
				onesNumbers = append(onesNumbers, number)
			} else {
				zeroNumbers = append(zeroNumbers, number)
			}
		}
		threshold := len(curNumbers) / 2
		if amountOfOnes > threshold || amountOfOnes*2 == len(curNumbers) {
			curNumbers = onesNumbers
		} else {
			curNumbers = zeroNumbers
		}
	}
	ogr, _ := strconv.ParseInt(curNumbers[0], 2, 0)

	curNumbers = numbers
	for curBit := 0; curBit < amountOfBits; curBit++ {
		if len(curNumbers) == 1 {
			break
		}
		amountOfOnes := 0

		onesNumbers := []string{}
		zeroNumbers := []string{}
		for _, number := range curNumbers {
			if number[curBit] == '1' {
				amountOfOnes++
				onesNumbers = append(onesNumbers, number)
			} else {
				zeroNumbers = append(zeroNumbers, number)
			}
		}
		threshold := len(curNumbers) / 2
		if amountOfOnes > threshold || amountOfOnes*2 == len(curNumbers) {
			curNumbers = zeroNumbers
		} else {
			curNumbers = onesNumbers
		}
	}
	co, _ := strconv.ParseInt(curNumbers[0], 2, 0)
	fmt.Println(ogr * co)
}
