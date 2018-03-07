/*
Package main executes a certain number of iterations of the game of the life.
*/
package main

import "github.com/jeremy-miller/life-go/internal/life"

const iterations = 5

func main() {
	life.Run(iterations)
}
