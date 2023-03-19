package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Words struct {
	words []string
}

func (w *Words) Lookup(word string) bool {
	left, right := 0, len(w.words)-1
	var middle int

	for left <= right {
		middle = (left + right) / 2
		if w.words[middle] == word {
			return true
		}
		if w.words[middle] < word {
			left = middle + 1
		}
		if w.words[middle] > word {
			right = middle - 1
		}
	}

	return false
}

func LoadWords(path string, minLength int) (*Words, error) {
	wordFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file %s: %w", path, err)
	}

	words := make([]string, 0)
	scanner := bufio.NewScanner(wordFile)

	for scanner.Scan() {
		word := scanner.Text()

		if len(word) < minLength {
			continue
		}

		word = strings.TrimSpace(word)
		word = strings.ToLower(word)

		words = append(words, word)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error when parsing word list: %w", err)
	}

	return &Words{words: words}, nil
}
