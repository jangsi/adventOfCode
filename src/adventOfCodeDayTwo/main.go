package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Rock     = "Rock"
	Paper    = "Paper"
	Scissors = "Scissors"

	LostScore = 0
	DrawScore = 3
	WinScore  = 6

	RockScore     = 1
	PaperScore    = 2
	ScissorsScore = 3
)

var codec = map[string]string{
	"A": Rock,
	"X": Rock,
	"B": Paper,
	"Y": Paper,
	"C": Scissors,
	"Z": Scissors,
}

var scoreByHand = map[string]int{
	Rock:     RockScore,
	Paper:    PaperScore,
	Scissors: ScissorsScore,
}

type game struct {
	self string
	beat string
}

var gameRounds = []game{
	{
		self: Rock,
		beat: Scissors,
	},
	{
		self: Paper,
		beat: Rock,
	},
	{
		self: Scissors,
		beat: Paper,
	},
}

var gameRoundByHand = map[string]game{
	Rock: {
		self: Rock,
		beat: Scissors,
	},
	Paper: {
		self: Paper,
		beat: Rock,
	},
	Scissors: {
		self: Scissors,
		beat: Paper,
	},
}

func main() {
	answerOne, err := parseStrategy(playRoundOne)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Advent day two part one: %d\n", answerOne)

	answerTwo, err := parseStrategy(playRoundReal)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("Advent day two part two: %d\n", answerTwo)
}

func parseStrategy(strat func(them string, you string) int) (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return -1, err
	}
	input := bufio.NewScanner(f)

	var total int
	for input.Scan() {
		text := input.Text()
		round := strings.Split(text, " ")
		score := strat(round[0], round[1])
		total += score
	}
	f.Close()
	return total, nil
}

// The winner of the whole tournament is the player with the highest score.
// Your total score is the sum of your scores for each round.
// The score for a single round is the score for the shape you selected
// (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round
// (0 if you lost, 3 if the round was a draw, and 6 if you won).
func playRoundOne(them string, you string) int {
	if codec[them] == codec[you] {
		return scoreByHand[codec[you]] + DrawScore
	}
	round := gameRoundByHand[codec[you]]
	if round.beat == codec[them] {
		return scoreByHand[codec[you]] + WinScore
	}
	return scoreByHand[codec[you]] + LostScore
}

// "Anyway, the second column says how the round needs to end: X means you need to lose,
// Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
func playRoundReal(them string, you string) int {
	switch you {
	case "X":
		round := gameRoundByHand[codec[them]]
		return LostScore + scoreByHand[round.beat]
	case "Y":
		return DrawScore + scoreByHand[codec[them]]
	case "Z":
		round := gameRoundByHand[gameRoundByHand[codec[them]].beat]
		return WinScore + scoreByHand[round.beat]
	}

	return 0
}
