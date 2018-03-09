package life

import (
	"fmt"
	"sync"
)

type cell struct {
	x, y, alive int
}

type board struct {
	cells  [][]int
	width  int
	height int
}

func newBoard(i *initialLayout) *board {
	board := newEmptyBoard(i)
	board = initializeBoard(board, i)
	return board
}

func newEmptyBoard(i *initialLayout) *board {
	cells := make([][]int, i.height)
	for j := range cells {
		cells[j] = make([]int, i.width)
	}
	return &board{cells: cells, width: i.width, height: i.height}
}

func initializeBoard(b *board, i *initialLayout) *board {
	for _, p := range i.positions {
		b.cells[p.y][p.x] = 1
	}
	return b
}

func (b *board) tick() {
	var tickedCells = make(chan cell, b.width*b.height)
	var wg sync.WaitGroup
	for y, column := range b.cells {
		for x, alive := range column {
			cell := cell{x: x, y: y, alive: alive}
			wg.Add(1)
			go tickCell(cell, *b, tickedCells, &wg)
		}
	}
	wg.Wait()
	updateBoard(b, tickedCells)
}

func tickCell(c cell, b board, tickedCells chan cell, wg *sync.WaitGroup) {
	defer wg.Done()
	c.alive = aliveAfterTick(c, b)
	tickedCells <- c
}

func aliveAfterTick(c cell, b board) int {
	livingNeighbors := getLivingNeighbors(c, b)
	alive := 0
	if c.alive == 1 {
		switch {
		case livingNeighbors < 2:
			alive = 0
		case livingNeighbors == 2 || livingNeighbors == 3:
			alive = 1
		default:
			alive = 0
		}
	} else if livingNeighbors == 3 {
		alive = 1
	}
	return alive
}

func getLivingNeighbors(c cell, b board) int {
	livingNeighbors := 0
	for i := -1; i <= 1; i++ {
		if i >= 0 && i < b.height {
			for j := -1; j <= 1; j++ {
				if j >= 0 && j < b.width {
					livingNeighbors += b.cells[c.y+i][c.x+j]
				}
			}
		}
	}
	livingNeighbors -= b.cells[c.y][c.x] // don't count the current cell as a neighbor
	return livingNeighbors
}

func updateBoard(b *board, tickedCells chan cell) {
	for c := range tickedCells {
		b.cells[c.y][c.x] = c.alive
	}
}

func (b *board) print() {
	for h := 0; h < b.height; h++ {
		for w := 0; w < b.width; w++ {
			if b.cells[h][w] == 1 {
				fmt.Print('O')
			} else {
				fmt.Print(' ')
			}
		}
		fmt.Println()
	}
}
