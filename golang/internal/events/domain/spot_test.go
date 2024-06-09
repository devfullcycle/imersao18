package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSpot(t *testing.T) {
	event, _ := NewEvent("Concert", "Stadium", "Music Inc.", RatingLivre, time.Now().Add(24*time.Hour), 100, 50.0, "http://x.jpg", 1)
	spot, err := NewSpot(event, "A1")
	assert.Nil(t, err)
	assert.NotNil(t, spot)
	assert.Equal(t, "A1", spot.Name)
	assert.Equal(t, SpotStatusAvailable, spot.Status)
	assert.Equal(t, event.ID, spot.EventID)
	assert.NotEmpty(t, spot.ID)
}

func TestSpot_Validate(t *testing.T) {
	event, _ := NewEvent("Concert", "Stadium", "Music Inc.", RatingLivre, time.Now().Add(24*time.Hour), 100, 50.0, "http://x.jpg", 1)
	spot := &Spot{
		EventID: event.ID,
		Name:    "",
		Status:  SpotStatusAvailable,
	}

	err := spot.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "spot name is required", err.Error())

	spot.Name = "1A"
	err = spot.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "spot name must start with a letter", err.Error())

	spot.Name = "A"
	err = spot.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "spot name must be at least 2 characters long", err.Error())

	spot.Name = "A1"
	err = spot.Validate()
	assert.Nil(t, err)
}

func TestSpot_Reserve(t *testing.T) {
	event, _ := NewEvent("Concert", "Stadium", "Music Inc.", RatingLivre, time.Now().Add(24*time.Hour), 100, 50.0, "http://x.jpg", 1)
	spot, _ := NewSpot(event, "A1")
	err := spot.Reserve("ticket123")
	assert.Nil(t, err)
	assert.Equal(t, SpotStatusSold, spot.Status)
	assert.Equal(t, "ticket123", spot.TicketID)
}

func TestSpot_Reserve_AlreadyReserved(t *testing.T) {
	event, _ := NewEvent("Concert", "Stadium", "Music Inc.", RatingLivre, time.Now().Add(24*time.Hour), 100, 50.0, "http://x.jpg", 1)
	spot, _ := NewSpot(event, "A1")
	spot.Reserve("ticket123")
	err := spot.Reserve("ticket456")
	assert.NotNil(t, err)
	assert.Equal(t, ErrSpotAlreadyReserved, err)
	assert.Equal(t, "ticket123", spot.TicketID)
}
