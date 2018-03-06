package life

type cell struct {
	alive bool
	x     int
	y     int
}

func new(x, y int) cell {
	return cell{false, x, y}
}
