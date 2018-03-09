/*
Package life implements a basic version of Conway's Game of Life in Go.
*/
package life

import "time"

// Run initializes the game and then runs it the provided number of iterations, printing the
// state of the board between iterations.
func Run(iterations int) {
	board := newBoard(&blinker)
	board.print()
	for i := 0; i < iterations; i++ {
		board.tick()
		board.print()
		time.Sleep(5 * time.Second)
	}
}
