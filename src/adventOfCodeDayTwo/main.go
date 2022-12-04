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
	var score int
	switch codec[you] {
	case Rock:
		score = RockScore
	case Paper:
		score = PaperScore
	case Scissors:
		score = ScissorsScore
	}
	if codec[them] == codec[you] {
		return score + DrawScore
	}
	if codec[them] == Rock {
		if codec[you] == Paper {
			return score + WinScore
		}
		if codec[you] == Scissors {
			return score + LostScore
		}
	}
	if codec[them] == Paper {
		if codec[you] == Scissors {
			return score + WinScore
		}
		if codec[you] == Rock {
			return score + LostScore
		}
	}
	if codec[them] == Scissors {
		if codec[you] == Rock {
			return score + WinScore
		}
		if codec[you] == Paper {
			return score + LostScore
		}
	}
	return score
}

// "Anyway, the second column says how the round needs to end: X means you need to lose,
// Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
func playRoundReal(them string, you string) int {
	var score int
	switch you {
	case "X":
		score = LostScore
		if codec[them] == Rock {
			return score + ScissorsScore
		}
		if codec[them] == Paper {
			return score + RockScore
		}
		if codec[them] == Scissors {
			return score + PaperScore
		}
	case "Y":
		score = DrawScore
		if codec[them] == Rock {
			return score + RockScore
		}
		if codec[them] == Paper {
			return score + PaperScore
		}
		if codec[them] == Scissors {
			return score + ScissorsScore
		}
	case "Z":
		score = WinScore
		if codec[them] == Rock {
			return score + PaperScore
		}
		if codec[them] == Paper {
			return score + ScissorsScore
		}
		if codec[them] == Scissors {
			return score + RockScore
		}
	}

	return score
}
