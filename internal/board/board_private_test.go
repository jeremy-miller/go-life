package board

import (
	"testing"

	"github.com/jeremy-miller/life-go/internal/configuration"
	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	expected := &Board{
		cells:  [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
		width:  5,
		height: 5,
	}
	i := configuration.Blinker()
	assert.Equal(t, expected, NewBoard(i))
}

func TestTick(t *testing.T) {
	expected := &Board{
		cells:  [][]int{{0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}},
		width:  5,
		height: 5,
	}
	b := &Board{
		cells:  [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
		width:  5,
		height: 5,
	}
	b.Tick()
	assert.Equal(t, expected, b)
}

func TestAliveAfterTick(t *testing.T) {
	expected := &Board{
		cells:  [][]int{{0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 1}},
		width:  5,
		height: 5,
	}
	b := &Board{
		cells:  [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}},
		width:  5,
		height: 5,
	}
	b.Tick()
	assert.Equal(t, expected, b)
}
