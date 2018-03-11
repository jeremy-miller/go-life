/*
Package life implements Conway's Game of Life.

Reference: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
*/
package life

import "time"

// Run initializes the game board with an initial configuration, then runs the game the provided
// number of iterations.  The initial board state is printed, as well as the state between iterations.
//
// Currently the only supported initial configuration is "blinker", which is used by default.
func Run(iterations int) {
	board := newBoard(&blinker)
	board.print()
	for i := 0; i < iterations; i++ {
		time.Sleep(1 * time.Second)
		board.tick()
		board.print()
	}
}
