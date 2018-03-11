package life_test

import (
	"github.com/jeremy-miller/life-go/internal/life"
)

func ExampleRun() {
	iterations := 1
	life.Run(iterations)
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
