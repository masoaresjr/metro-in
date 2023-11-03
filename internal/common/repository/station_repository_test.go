package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewStationRepository test if StationRepository constructor is returning the correct implementation
func TestNewStationRepository(t *testing.T) {
	stationRepository := NewStationRepository(nil)

	// stationRepository can't be null
	assert.NotNil(t, stationRepository)

	// stationRepository must implement StationRepository interface
	assert.Implements(t, (*StationRepository)(nil), stationRepository)
}
