package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	answerOne, err := sumPriorities()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Advent day one part one: %d\n", answerOne)
}

func runeToPriority(r rune) int {
	val := int(r)
	if 64 < val && val < 91 {
		return int(r) - 64 + 26
	}
	if 96 < val && val < 123 {
		return int(r) - 96
	}
	return 0
}

func sumPriorities() (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	input := bufio.NewScanner(f)
	var corpus = make(map[rune]int)
	var priority int
	for input.Scan() {
		text := input.Text()
		halfway := len(text) / 2
		h1 := text[:halfway]
		h2 := text[halfway:]
		for _, c := range h1 {
			corpus[c] = 1
		}
		for _, c := range h2 {
			if _, ok := corpus[c]; ok {
				priority += runeToPriority(c)
				delete(corpus, c)
			}
		}
		corpus = make(map[rune]int)
	}
	f.Close()
	return priority, nil
}
