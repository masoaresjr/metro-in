package controller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewStationController test if SubwayLineController constructor is returning the correct implementation
func TestNewStationController(t *testing.T) {
	stationController := NewStationController(nil)

	// stationController can't be null
	assert.NotNil(t, stationController)

	// stationController must implement SubwayLineController interface
	assert.Implements(t, (*StationController)(nil), stationController)
}
