package board_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/jeremy-miller/life-go/internal/board"
	"github.com/jeremy-miller/life-go/internal/configuration"
	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	i := configuration.Blinker()
	assert.IsType(t, &board.Board{}, board.NewBoard(i))
}

func TestTick(t *testing.T) {
	i := configuration.Blinker()
	b := board.NewBoard(i)
	assert.NotPanics(t, func() { b.Tick() })
}

func TestPrint(t *testing.T) {
	expected := `.....
.....
.OOO.
.....
.....`
	rescueStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w
	i := configuration.Blinker()
	b := board.NewBoard(i)
	b.Print() // gets captured
	if err := w.Close(); err != nil {
		t.Error(err)
	}
	out, _ := ioutil.ReadAll(r)
	os.Stderr = rescueStderr
	assert.Contains(t, expected, string(out))
}
