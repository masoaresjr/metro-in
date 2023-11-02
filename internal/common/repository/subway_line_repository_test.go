package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewSubwayLineController test if SubwayLineController constructor is returning the correct implementation
func TestNewSubwayLineRepository(t *testing.T) {
	subwayLineRepository := NewSubwayLineRepository(nil)

	// subwayLineController can't be null
	assert.NotNil(t, subwayLineRepository)

	// subwayLineController must implement SubwayLineController interface
	assert.Implements(t, (*SubwayLineRepository)(nil), subwayLineRepository)
}