/*
Package game executes Conway's Game of Life.

Reference: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
*/
package game

import (
	"time"

	"github.com/jeremy-miller/life-go/internal/board"
	"github.com/jeremy-miller/life-go/internal/configuration"
)

/*
Run initializes the game board with an initial configuration, then runs the game
a certain number of iterations.  The initial board state is printed, as well as
the board state between iterations.

Currently the only supported initial configuration is "blinker", which is used
by default.
*/
func Run(iterations int) {
	config := configuration.Blinker()
	b := board.NewBoard(config)
	b.Print()
	for i := 0; i < iterations; i++ {
		time.Sleep(1 * time.Second) // to user time to see the board changes
		b.Tick()
		b.Print()
	}
}
