package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewEvent(t *testing.T) {
	event, err := NewEvent("Concert", "Stadium", "Music Inc.", RatingLivre, time.Now().Add(24*time.Hour), 100, 50.0, "http://x.jpg", 1)
	assert.Nil(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, "Concert", event.Name)
	assert.Equal(t, "Stadium", event.Location)
	assert.Equal(t, "Music Inc.", event.Organization)
	assert.Equal(t, RatingLivre, event.Rating)
	assert.Equal(t, 100, event.Capacity)
	assert.Equal(t, 50.0, event.Price)
	assert.Equal(t, 1, event.PartnerID)
	assert.NotEmpty(t, event.ID)
	assert.Empty(t, event.Spots)
	assert.Empty(t, event.Tickets)
}

func TestEvent_Validate(t *testing.T) {
	event := &Event{
		Name:     "",
		Date:     time.Now().Add(24 * time.Hour),
		Capacity: 100,
		Price:    50.0,
	}

	err := event.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "event name is required", err.Error())

	event.Name = "Concert"
	event.Date = time.Now().Add(-24 * time.Hour)
	err = event.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "event date must be in the future", err.Error())

	event.Date = time.Now().Add(24 * time.Hour)
	event.Capacity = -1
	err = event.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "event capacity must be greater than zero", err.Error())

	event.Capacity = 100
	event.Price = -10
	err = event.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "event price must be greater than zero", err.Error())
}

func TestEvent_AddSpot(t *testing.T) {
	event, err := NewEvent("Concert", "Stadium", "Music Inc.", RatingLivre, time.Now().Add(24*time.Hour), 100, 50.0, "http://x.jpg", 1)
	assert.Nil(t, err)
	assert.NotNil(t, event)

	spot, err := event.AddSpot("A1")
	assert.Nil(t, err)
	assert.NotNil(t, spot)
	assert.Equal(t, "A1", spot.Name)
	assert.Equal(t, SpotStatusAvailable, spot.Status)
	assert.Equal(t, event.ID, spot.EventID)
	assert.Equal(t, 1, len(event.Spots))
}
