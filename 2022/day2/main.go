package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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

func part2(bytesStream []byte) {
	total := 0
	scanner := bufio.NewScanner(bytes.NewReader(bytesStream))
	for scanner.Scan() {
		opAndEnd := strings.Split(scanner.Text(), " ")
		op := opAndEnd[0]
		end := opAndEnd[1]
		total += calcStrategy(op, end)
	}

	fmt.Println(total)
}

func calcStrategy(opponent, end string) int {
	score := 0
	switch end {
	case "X":
		score = lose
	case "Y":
		score = draw
	case "Z":
		score = win
	}

	if opponent == "A" {
		if score == lose {
			score += scissors
		} else if score == draw {
			score += rock
		} else if score == win {
			score += paper
		}
	} else if opponent == "B" {
		if score == lose {
			score += rock
		} else if score == draw {
			score += paper
		} else if score == win {
			score += scissors
		}
	} else if opponent == "C" {
		if score == lose {
			score += paper
		} else if score == draw {
			score += scissors
		} else if score == win {
			score += rock
		}
	}

	return score
}

const (
	rock     = 1
	paper    = 2
	scissors = 3

	lose = 0
	draw = 3
	win  = 6
)

func calcRoundScore(opponent, you string) int {
	score := 0
	switch you {
	case "X":
		score = rock
	case "Y":
		score = paper
	case "Z":
		score = scissors
	}

	if opponent == "A" {
		if score == rock {
			score += draw
		} else if score == paper {
			score += win
		} else if score == scissors {
			score += lose
		}
	} else if opponent == "B" {
		if score == rock {
			score += lose
		} else if score == paper {
			score += draw
		} else if score == scissors {
			score += win
		}
	} else if opponent == "C" {
		if score == rock {
			score += win
		} else if score == paper {
			score += lose
		} else if score == scissors {
			score += draw
		}
	}

	return score
}

func part1(bytesStream []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(bytesStream))
	totalScore := 0
	for scanner.Scan() {
		opAndYou := strings.Split(scanner.Text(), " ")
		op := opAndYou[0]
		you := opAndYou[1]
		totalScore += calcRoundScore(op, you)
	}
	fmt.Println(totalScore)
}
