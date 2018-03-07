package life

type position struct {
	x, y int
}

type initialLayout struct {
	positions []position
	width     int
	height    int
}

var blinker = initialLayout{
	positions: []position{position{1, 2}, position{2, 2}, position{3, 2}},
	width:     5,
	height:    5,
}
