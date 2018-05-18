/*
Package configuration contains configuration functionality for the Game of Life.
*/
package configuration

/*
Position defines one pair of X and Y coordinates within the game board.
*/
type Position struct {
	X, Y int
}

/*
InitialLayout defines the inital layout of the game board.
*/
type InitialLayout struct {
	Positions []Position // initial positions which are "alive"
	Width     int        // width of board in "cells"
	Height    int        // height of board in "cells"
}

/*
Blinker returns an initial layout of the "blinker" oscillator.

Reference: http://conwaylife.com/w/index.php?title=Blinker
*/
func Blinker() *InitialLayout {
	return &InitialLayout{
		Positions: []Position{{1, 2}, {2, 2}, {3, 2}},
		Width:     5,
		Height:    5,
	}
}
