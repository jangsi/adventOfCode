package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	answerOne, err := sumPriorities(basicPriorities)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Advent day three part one: %d\n", answerOne)

	answerTwo, err := sumPriorities(findBadgeRunePriorities)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("Advent day three part two: %d\n", answerTwo)
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

func findBadgeRunePriorities(group []string) (int, []string) {
	if len(group) < 3 {
		return 0, group
	}
	var corpus = make(map[rune]int)
	var candidates = make(map[rune]int)
	var priority int
	for _, text := range group {
		halfway := len(text) / 2
		h1 := text[:halfway]
		h2 := text[halfway:]
		for _, c := range h1 {
			corpus[c] = 1
		}
		for _, c := range h2 {
			if _, ok := corpus[c]; !ok {
				corpus[c] = -1
			} else {
				if corpus[c] != -1 {
					corpus[c] += 1
				}
			}
		}
		for r, occurence := range corpus {
			if occurence == 1 || occurence == -1 {
				candidates[r] -= 1
			}
		}
		corpus = make(map[rune]int)
	}
	for r, occurence := range candidates {
		if occurence == -3 {
			priority += runeToPriority(r)
		}
	}
	return priority, []string{}
}

func basicPriorities(group []string) (int, []string) {
	text := group[0]
	var corpus = make(map[rune]int)
	var priority int
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
	group = group[len(group):]
	return priority, group
}

func sumPriorities(alg func(group []string) (int, []string)) (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	input := bufio.NewScanner(f)

	var priority int
	group := []string{}
	for input.Scan() {
		text := input.Text()
		group = append(group, text)
		p, g := alg(group)
		priority += p
		group = g
	}
	f.Close()
	return priority, nil
}
