package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	filename := flag.String("input", "test_input.txt", "path to the data")
	flag.Parse()

	input, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("cannot open %s: %v\n", *filename, err)
	}
	part1(input)
	if err := input.Close(); err != nil {
		log.Println(err)
	}
}

func part1(input io.ReadCloser) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	numbers := strings.Split(scanner.Text(), ",")
	fmt.Println(numbers)
}
