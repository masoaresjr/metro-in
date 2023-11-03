package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewStationService test if StationService constructor is returning the correct implementation
func TestNewStationService(t *testing.T) {
	stationService := NewStationService(nil)

	// stationService can't be null
	assert.NotNil(t, stationService)

	// stationService must implement StationService interface
	assert.Implements(t, (*StationService)(nil), stationService)
}
