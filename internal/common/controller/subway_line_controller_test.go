package controller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewSubwayLineController test if SubwayLineController constructor is returning the correct implementation
func TestNewSubwayLineController(t *testing.T) {
	subwayLineController := NewSubwayLineController(nil)

	// subwayLineController can't be null
	assert.NotNil(t, subwayLineController)

	// subwayLineController must implement SubwayLineController interface
	assert.Implements(t, (*SubwayLineController)(nil), subwayLineController)
}
