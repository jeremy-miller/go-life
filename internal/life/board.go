package life

type board struct {
	cells  [][]bool
	width  int
	height int
}

func newBoard(i *initialLayout) *board {
	board := newEmptyBoard(i)
	board = initializeBoard(board, i)
	return board
}

func newEmptyBoard(i *initialLayout) *board {
	cells := make([][]bool, i.width)
	for j := range cells {
		cells[j] = make([]bool, i.height)
	}
	return &board{cells: cells, width: i.width, height: i.height}
}

func initializeBoard(b *board, i *initialLayout) *board {
	for _, p := range i.positions {
		b.cells[p.x][p.y] = true
	}
	return b
}

func (b *board) tick() {
	// iterate over each cell using goroutines
	// check if alive/dead, update cell accordingly
	// NOTE: don't update pointer to board
}
