/*
Package board implements the board of the Game of Life.
*/
package board

import (
	"fmt"
	"sync"

	"github.com/jeremy-miller/life-go/internal/configuration"
)

/*
Board defines a Game of Life board.
*/
type Board struct {
	cells  [][]int
	width  int
	height int
}

/*
NewBoard initializes and returns a new Game of Life board based on the provided
initial configuration.
*/
func NewBoard(i *configuration.InitialLayout) *Board {
	board := emptyBoard(i.Width, i.Height)
	board = board.initialize(i.Positions)
	return board
}

func emptyBoard(width, height int) *Board {
	cells := make([][]int, height)
	for j := range cells {
		cells[j] = make([]int, width)
	}
	return &Board{cells: cells, width: width, height: height}
}

func (b *Board) initialize(positions []configuration.Position) *Board {
	for _, p := range positions {
		b.cells[p.Y][p.X] = 1
	}
	return b
}

/*
Tick executes a "step" in time of the board.
*/
func (b *Board) Tick() {
	var wg sync.WaitGroup
	var tickedCells = make(chan cell, b.width*b.height)
	for y, column := range b.cells {
		for x, alive := range column {
			c := cell{x: x, y: y, alive: alive}
			wg.Add(1)
			go c.tick(*b, tickedCells, &wg)
		}
	}
	wg.Wait()
	close(tickedCells)
	b.update(tickedCells)
}

func (b *Board) update(tickedCells chan cell) {
	for c := range tickedCells {
		b.cells[c.y][c.x] = c.alive
	}
}

/*
Print prints the current state of the board.
*/
func (b *Board) Print() {
	for h := 0; h < b.height; h++ {
		for w := 0; w < b.width; w++ {
			if b.cells[h][w] == 1 {
				fmt.Print("O")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type cell struct {
	x, y, alive int
}

func (c *cell) tick(b Board, tickedCells chan cell, wg *sync.WaitGroup) {
	defer wg.Done()
	c.alive = c.aliveAfterTick(b)
	tickedCells <- *c
}

func (c *cell) aliveAfterTick(b Board) int {
	livingNeighbors := c.livingNeighbors(b)
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

func (c *cell) livingNeighbors(b Board) int {
	livingNeighbors := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			neighborX := c.x + j
			neighborY := c.y + i
			if neighborX >= 0 && neighborY >= 0 && neighborX < b.width && neighborY < b.height {
				livingNeighbors += b.cells[c.y+i][c.x+j]
			}
		}
	}
	livingNeighbors -= b.cells[c.y][c.x] // don't count the current cell as a neighbor
	return livingNeighbors
}
