package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	answerOne, err := countHighestNCalories(1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Advent day one part one: %d\n", answerOne)

	answerTwo, err := countHighestNCalories(3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Advent day one part two: %d\n", answerTwo)
}

func countHighestNCalories(n int) (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return -1, err
	}
	input := bufio.NewScanner(f)
	var topN nNumbers
	topN.n = n
	var currentElf int
	for input.Scan() {
		text := input.Text()
		if text == "" { // new elf
			keepTopN(&topN, currentElf)
			currentElf = 0
			continue
		}
		i, err := strconv.Atoi(text)
		if err != nil {
			return -1, err
		}
		currentElf += i
	}
	f.Close()
	var total int
	for _, number := range topN.numbers {
		total += number
	}
	return total, nil
}

type nNumbers struct {
	numbers []int
	n       int
}

func keepTopN(currentNNumbers *nNumbers, contender int) {
	currentNNumbers.numbers = append(currentNNumbers.numbers, contender)
	sort.Slice(currentNNumbers.numbers, func(i, j int) bool {
		return currentNNumbers.numbers[i] > currentNNumbers.numbers[j] // descending
	})
	var keep int
	if currentNNumbers.n > len(currentNNumbers.numbers) {
		keep = len(currentNNumbers.numbers)
	} else {
		keep = currentNNumbers.n
	}
	currentNNumbers.numbers = currentNNumbers.numbers[:keep]
}
