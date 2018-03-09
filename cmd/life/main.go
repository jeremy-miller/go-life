/*
Package main executes a certain number of iterations of the game of the life.
*/
package main

import (
	"os"
	"strconv"

	"github.com/jeremy-miller/life-go/internal/life"
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
	life.Run(iterations)
}
