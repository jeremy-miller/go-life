package life

type board struct {
	cells  [][]bool
	width  int
	height int
}

func newBoard(width, height int, initial []cell) *board {
	board := newEmptyBoard(width, height)
	board = initializeBoard(board, initial)
	return board
}

func newEmptyBoard(width, height int) *board {
	cells := make([][]bool, width)
	for i := range cells {
		cells[i] = make([]bool, height)
	}
	return &board{cells: cells, width: width, height: height}
}

func initializeBoard(b *board, cells []cell) *board {
	for c := range cells {
		b[c.x][x.y] = true
	}
	return b
}

func (b *board) tick() {
	// iterate over each cell using goroutines
	// check if alive/dead, update cell accordingly
	// NOTE: don't update pointer to board
}
