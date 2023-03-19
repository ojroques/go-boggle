package main

import (
	"fmt"
	"sort"
	"strings"
)

type Results struct {
	results map[string]struct{}
	words   []string
}

func NewResults() *Results {
	return &Results{results: make(map[string]struct{})}
}

func (r *Results) Add(word string) {
	r.results[word] = struct{}{}
}

func (r *Results) Length() int {
	return len(r.results)
}

func (r *Results) Slice() []string {
	if r.words != nil {
		return r.words
	}

	r.words = make([]string, 0, len(r.results))

	for word := range r.results {
		r.words = append(r.words, word)
	}

	sort.Slice(r.words, func(i, j int) bool {
		if len(r.words[i]) == len(r.words[j]) {
			return r.words[i] < r.words[j]
		}
		return len(r.words[i]) < len(r.words[j])
	})

	return r.words
}

func (r *Results) Points() int {
	points := 0

	for _, word := range r.Slice() {
		if len(word) < 5 {
			points += 1
		}
		if len(word) == 5 {
			points += 2
		}
		if len(word) == 6 {
			points += 3
		}
		if len(word) == 7 {
			points += 5
		}
		if len(word) > 7 {
			points += 11
		}
	}

	return points
}

func (r *Results) String() string {
	var sb strings.Builder

	for _, word := range r.Slice() {
		sb.WriteString(fmt.Sprintln(word))
	}

	return sb.String()
}
