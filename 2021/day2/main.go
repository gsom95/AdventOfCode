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
	defer input.Close()

	input, err = os.Open(*filename)
	if err != nil {
		log.Fatalf("cannot open %s: %v\n", *filename, err)
	}
	part2(input)
	defer input.Close()
}

func part1(input io.ReadCloser) {
	const (
		forwardString string = "forward"
		downString    string = "down"
		upString      string = "up"
	)

	vertical := 0
	horizontal := 0

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()
		textSplit := strings.Split(text, " ")
		cmd := textSplit[0]
		delta, _ := strconv.Atoi(textSplit[1])

		switch cmd {
		case forwardString:
			horizontal += delta
		case downString:
			vertical += delta
		case upString:
			vertical -= delta
		}

	}

	fmt.Println(vertical * horizontal)
}

func part2(input io.ReadCloser) {
	const (
		forwardString string = "forward"
		downString    string = "down"
		upString      string = "up"
	)

	vertical := 0
	horizontal := 0
	aim := 0

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()
		textSplit := strings.Split(text, " ")
		cmd := textSplit[0]
		delta, _ := strconv.Atoi(textSplit[1])

		switch cmd {
		case forwardString:
			horizontal += delta
			vertical += (delta * aim)
		case downString:
			aim += delta
		case upString:
			aim -= delta
		}

	}

	fmt.Println(vertical * horizontal)
}
