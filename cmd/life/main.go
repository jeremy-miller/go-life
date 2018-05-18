/*
Package main executes Conway's Game of Life.

An integer can be passed on the command line to control the number of iterations
of the game (default 5).
*/
package main

import (
	"os"
	"strconv"

	"github.com/jeremy-miller/life-go/internal/game"
)

func main() {
	iterations := 5
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic("Iteration value must be a number")
		}
		iterations = n
	}
	game.Run(iterations)
}
