package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewSubwayLineService test if SubwayLineService constructor is returning the correct implementation
func TestNewSubwayLineService(t *testing.T) {
	subwayLineService := NewSubwayLineService(nil)

	// subwayLineController can't be null
	assert.NotNil(t, subwayLineService)

	// subwayLineController must implement SubwayLineController interface
	assert.Implements(t, (*SubwayLineService)(nil), subwayLineService)
}
