package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

var isTest = false

func main() {
	filename := "input.txt"
	if isTest {
		filename = "test_" + filename
	}

	input, err := os.Open(filename)
	if err != nil {
		log.Fatalf("cannot open the file '%s': %v", filename, err)
	}
	calories := calcCalories(input)
	if err = input.Close(); err != nil {
		log.Fatalln("cannot close the file:", err)
	}

	fmt.Println("part1:", calories[0])
	fmt.Println("part2:", calories[0]+calories[1]+calories[2])
}

func calcCalories(input io.ReadCloser) []int {
	scanner := bufio.NewScanner(input)

	cals := []int{}
	curCal := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			cals = append(cals, curCal)
			curCal = 0
			continue
		}

		cal, err := strconv.Atoi(text)
		if err != nil {
			log.Fatalln("cannot parse number:", err)
		}
		curCal += cal
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cals)))
	return cals
}
