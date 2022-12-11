package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	answerOne, err := countEclipsedPairs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Advent day four part one: %d\n", answerOne)

	answerTwo, err := countOverlappingPairs()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("Advent day four part two: %d\n", answerTwo)
}

func parseSections(text string) ([]int, []int) {
	sections := strings.Split(text, ",")
	sOne := sections[0]
	sTwo := sections[1]
	rOne := strings.Split(sOne, "-")
	lOne, _ := strconv.Atoi(rOne[0])
	hOne, _ := strconv.Atoi(rOne[1])
	rangeOne := []int{lOne, hOne}
	rTwo := strings.Split(sTwo, "-")
	lTwo, _ := strconv.Atoi(rTwo[0])
	hTwo, _ := strconv.Atoi(rTwo[1])
	rangeTwo := []int{lTwo, hTwo}
	return rangeOne, rangeTwo
}

func countOverlappingPairs() (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	input := bufio.NewScanner(f)

	var count int
	for input.Scan() {
		text := input.Text()
		rangeOne, rangeTwo := parseSections(text)
		var overlappingRange []int
		if rangeOne[0] <= rangeTwo[0] && rangeTwo[0] <= rangeOne[1] {
			overlappingRange = rangeOne
		}
		if rangeOne[0] <= rangeTwo[1] && rangeTwo[1] <= rangeOne[1] {
			overlappingRange = rangeOne
		}
		if rangeTwo[0] <= rangeOne[0] && rangeOne[0] <= rangeTwo[1] {
			overlappingRange = rangeTwo
		}
		if rangeTwo[0] <= rangeOne[1] && rangeOne[1] <= rangeTwo[1] {
			overlappingRange = rangeTwo
		}
		if len(overlappingRange) > 0 {
			count++
		}
	}
	f.Close()
	return count, nil
}

func countEclipsedPairs() (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	input := bufio.NewScanner(f)

	var count int
	for input.Scan() {
		text := input.Text()
		rangeOne, rangeTwo := parseSections(text)
		var containingRange []int
		if rangeOne[0] <= rangeTwo[0] && rangeTwo[1] <= rangeOne[1] {
			containingRange = rangeOne
		}
		if rangeTwo[0] <= rangeOne[0] && rangeOne[1] <= rangeTwo[1] {
			containingRange = rangeTwo
		}
		if len(containingRange) > 0 {
			count++
		}
	}
	f.Close()
	return count, nil
}
