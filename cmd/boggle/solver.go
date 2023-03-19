package main

import (
	"strings"
)

type Solver struct {
	boggle   *Boggle
	words    *Words
	visited  map[Position]bool
	startPos Position
	out      chan<- string
}

func NewSolver(boggle *Boggle, words *Words, startPos Position) *Solver {
	return &Solver{
		boggle:   boggle,
		words:    words,
		startPos: startPos,
	}
}

func (s *Solver) Run(out chan<- string) {
	s.out = out
	s.visited = make(map[Position]bool)
	s.run(s.startPos, make([]Character, 0))
}

func (s *Solver) check(currentWord []Character) {
	var sb strings.Builder

	for _, character := range currentWord {
		sb.WriteString(character.String())
	}

	if s.words.Lookup(sb.String()) {
		s.out <- sb.String()
	}
}

func (s *Solver) run(currentPos Position, currentWord []Character) {
	s.visited[currentPos] = true

	currentWord = append(currentWord, s.boggle.Character(currentPos))
	s.check(currentWord)

	for _, neighbor := range s.boggle.Neighbours(currentPos) {
		if visited, ok := s.visited[neighbor]; ok && visited {
			continue
		}

		s.run(neighbor, currentWord)
	}

	currentWord = currentWord[:len(currentWord)-1]
	s.visited[currentPos] = false
}
