package main

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
	s.run(s.startPos, "")
}

func (s *Solver) run(currentPos Position, currentWord string) {
	s.visited[currentPos] = true

	currentWord = currentWord + string(s.boggle.Character(currentPos))
	if s.words.Lookup(currentWord) {
		s.out <- currentWord
	}

	for _, neighbor := range s.boggle.Neighbours(currentPos) {
		if visited, ok := s.visited[neighbor]; ok && visited {
			continue
		}

		s.run(neighbor, currentWord)
	}

	currentWord = currentWord[:len(currentWord)-1]
	s.visited[currentPos] = false
}
