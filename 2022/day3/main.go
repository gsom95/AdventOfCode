package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
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

type found struct{}
type set map[rune]found

func priority(char rune) int {
	value := 1
	if unicode.IsUpper(char) {
		value += 26

	}
	char = unicode.ToUpper(char)
	value += int(char) % 65
	return value
}

func part1(bytesStream []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(bytesStream))
	scanner.Split(bufio.ScanLines)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)

		items := set{}
		for _, item := range line[:lineLen/2] {
			items[item] = found{}
		}
		for _, item := range line[lineLen/2:] {
			if _, itemFound := items[item]; itemFound {
				total += priority(item)
				break
			}
		}
	}

	fmt.Println(total)
}

func part2(bytesStream []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(bytesStream))
	scanner.Split(bufio.ScanLines)

	total := 0

	lines := make([]string, 0, 3)
	for scanner.Scan() {
		line := scanner.Text()
		if len(lines) < 3 {
			lines = append(lines, line)
		}
		if len(lines) < 3 {
			continue
		}

		items := set{}
		for _, item := range lines[0] {
			items[item] = found{}
		}
		commonItems := set{}
		for _, item := range lines[1] {
			_, itemFound := items[item]
			if itemFound {
				commonItems[item] = found{}
			}
		}
		for _, item := range lines[2] {
			if _, itemFound := commonItems[item]; itemFound {
				total += priority(item)
				break
			}
		}

		lines = make([]string, 0, 3)
	}

	fmt.Println(total)
}
