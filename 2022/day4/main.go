package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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
	bytesRead, err := io.ReadAll(input)
	if err != nil {
		log.Fatalf("cannot read the file '%s': %v", filename, err)
	}

	part1(bytesRead)
	part2(bytesRead)
}

type segment [2]int

func part1(bytesStream []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(bytesStream))
	scanner.Split(bufio.ScanLines)

	overlaps := 0
	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, ",")

		segments := [2]segment{}
		for i, pair := range pairs {
			limits := strings.Split(pair, "-")
			segments[i][0], _ = strconv.Atoi(limits[0])
			segments[i][1], _ = strconv.Atoi(limits[1])
		}
		if (segments[0][0] >= segments[1][0] && segments[0][1] <= segments[1][1]) ||
			(segments[1][0] >= segments[0][0] && segments[1][1] <= segments[0][1]) {
			overlaps++
			continue
		}
	}

	fmt.Println(overlaps)
}

func part2(bytesStream []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(bytesStream))
	scanner.Split(bufio.ScanLines)

	overlaps := 0
	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, ",")

		segments := [2]segment{}
		for i, pair := range pairs {
			limits := strings.Split(pair, "-")
			segments[i][0], _ = strconv.Atoi(limits[0])
			segments[i][1], _ = strconv.Atoi(limits[1])
		}
		if segments[0][1] < segments[1][0] || segments[1][1] < segments[0][0] {
			continue
		}
		overlaps++
	}

	fmt.Println(overlaps)
}
