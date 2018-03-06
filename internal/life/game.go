/*
Package life implements a basic version of Conway's Game of Life in Go.
*/
package life

const boardWidth = 5
const boardHeight = 6

// Run initializes the game and then runs it the provided number of iterations.
func Run(iterations int) {
	board := newBoard(boardWidth, boardHeight, blinker)
	for i := 0; i < iterations; i++ {
		board.tick()
	}
}
