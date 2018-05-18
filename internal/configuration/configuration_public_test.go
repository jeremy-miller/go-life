package configuration_test

import (
	"testing"

	"github.com/jeremy-miller/life-go/internal/configuration"
	"github.com/stretchr/testify/assert"
)

func TestBlinker(t *testing.T) {
	expected := &configuration.InitialLayout{
		Positions: []configuration.Position{{X: 1, Y: 2}, {X: 2, Y: 2}, {X: 3, Y: 2}},
		Width:     5,
		Height:    5,
	}
	assert.Equal(t, expected, configuration.Blinker())
}
