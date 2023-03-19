package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Character struct {
	character rune
}

type Position struct {
	line   int
	column int
}

type Boggle struct {
	grid [4][4]Character
}

func NewCharacter(character rune) Character {
	return Character{character: unicode.ToLower(character)}
}

func (c Character) String() string {
	return fmt.Sprintf("%c", c.character)
}

func NewPosition(line, column int) Position {
	return Position{line: line, column: column}
}

func (p Position) Line() int {
	return p.line
}

func (p Position) Column() int {
	return p.column
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.Line(), p.Column())
}

func NewBoggle(grid [4][4]Character) *Boggle {
	return &Boggle{grid: grid}
}

func NewBoggleFromFile(path string) (*Boggle, error) {
	gridFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file %s: %w", path, err)
	}

	grid := [4][4]Character{}
	scanner := bufio.NewScanner(gridFile)

	l := 0
	for scanner.Scan() {
		for c, character := range strings.TrimSpace(scanner.Text()) {
			grid[l][c] = NewCharacter(character)
		}
		l += 1
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error when parsing boggle grid: %w", err)
	}

	return NewBoggle(grid), nil
}

func (b *Boggle) Character(position Position) Character {
	return b.grid[position.Line()][position.Column()]
}

func (b *Boggle) Neighbours(position Position) []Position {
	positions := make([]Position, 0)

	for _, l := range [3]int{-1, 0, 1} {
		for _, c := range [3]int{-1, 0, 1} {
			currentLine, currentColumn := position.Line()+l, position.Column()+c
			if currentLine == position.Line() && currentColumn == position.Column() {
				continue
			}
			if currentLine < 0 || currentLine > 3 {
				continue
			}
			if currentColumn < 0 || currentColumn > 3 {
				continue
			}
			positions = append(positions, NewPosition(currentLine, currentColumn))
		}
	}

	return positions
}

func (b *Boggle) String() string {
	var sb strings.Builder

	for l := range b.grid {
		for c := range b.grid[l] {
			sb.WriteString(b.grid[l][c].String())
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
