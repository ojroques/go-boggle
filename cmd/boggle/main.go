package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
)

func panicOnErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func buildWords(wordsPath string) *Words {
	words, err := LoadWords(wordsPath, 3)
	panicOnErr(err)
	return words
}

func buildBoggle(gridPath string) *Boggle {
	boggle, err := NewBoggleFromFile(gridPath)
	panicOnErr(err)
	return boggle
}

func main() {
	// Config
	wordsPath := flag.String("w", "words.txt", "word list")
	gridPath := flag.String("b", "grid.txt", "boggle grid")
	flag.Parse()

	// Variables
	words := buildWords(*wordsPath)
	boggle := buildBoggle(*gridPath)
	results := NewResults()
	var wg sync.WaitGroup
	wordChan := make(chan string)

	// Run Solvers
	for l := 0; l < 4; l++ {
		for c := 0; c < 4; c++ {
			wg.Add(1)
			go func(line, column int) {
				defer wg.Done()
				solver := NewSolver(boggle, words, NewPosition(line, column))
				solver.Run(wordChan)
			}(l, c)
		}
	}

	// Wait for Solvers to finish
	go func() {
		wg.Wait()
		close(wordChan)
	}()

	// Process results
	for word := range wordChan {
		results.Add(word)
	}

	// Print results
	fmt.Printf("%d words found:\n", results.Length())
	fmt.Print(results)
	fmt.Printf("%d points\n", results.Points())
}
