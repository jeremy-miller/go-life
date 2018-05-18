package game_test

import (
	"github.com/jeremy-miller/life-go/internal/game"
)

func ExampleRun() {
	iterations := 1
	game.Run(iterations)
	// Output: .....
	//.....
	//.OOO.
	//.....
	//.....
	//
	//.....
	//..O..
	//..O..
	//..O..
	//.....
}
