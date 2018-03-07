/*
Package life implements a basic version of Conway's Game of Life in Go.
*/
package life

// Run initializes the game and then runs it the provided number of iterations.
func Run(iterations int) {
	board := newBoard(&blinker)
	for i := 0; i < iterations; i++ {
		board.tick()
	}
}
