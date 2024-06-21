package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTicket(t *testing.T) {
	event, _ := NewEvent("Concert", "Stadium", "Music Inc.", RatingLivre, time.Now().Add(24*time.Hour), 100, 50.0, "http://x.jpg", 1)
	spot, _ := NewSpot(event, "A1")
	ticket, err := NewTicket(event, spot, TicketKindFull)
	assert.Nil(t, err)
	assert.NotNil(t, ticket)
	assert.Equal(t, TicketKindFull, ticket.TicketKind)
	assert.Equal(t, 50.0, ticket.Price)
	assert.Equal(t, event.ID, ticket.EventID)
	assert.Equal(t, spot.ID, ticket.Spot.ID)
	assert.NotEmpty(t, ticket.ID)
}

func TestNewTicket_HalfPrice(t *testing.T) {
	event, _ := NewEvent("Concert", "Stadium", "Music Inc.", RatingLivre, time.Now().Add(24*time.Hour), 100, 50.0, "http://x.jpg", 1)
	spot, _ := NewSpot(event, "A1")
	ticket, err := NewTicket(event, spot, TicketKindHalf)
	assert.Nil(t, err)
	assert.NotNil(t, ticket)
	assert.Equal(t, TicketKindHalf, ticket.TicketKind)
	assert.Equal(t, 25.0, ticket.Price)
}

func TestTicket_Validate(t *testing.T) {
	ticket := &Ticket{
		Price: -10,
	}

	err := ticket.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "ticket price must be greater than zero", err.Error())
}
