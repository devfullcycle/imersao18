package domain

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidTicketType = errors.New("invalid ticket type")

// TicketType represents the type of a ticket.
type TicketType string

const (
	TicketTypeHalf TicketType = "half" // Half-price ticket
	TicketTypeFull TicketType = "full" // Full-price ticket
)

// Ticket represents a ticket for an event.
type Ticket struct {
	ID            string
	EventID       string
	Spot          *Spot
	TicketType    TicketType
	Price         float64
	ReservationID int
}

// IsValidTicketType checks if a ticket type is valid.
func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

// NewTicket creates a new ticket with the given parameters.
func NewTicket(event *Event, spot *Spot, ticketType TicketType) (*Ticket, error) {
	if !IsValidTicketType(ticketType) {
		return nil, ErrInvalidTicketType
	}

	price := event.Price
	if ticketType == TicketTypeHalf {
		price = price / 2
	}

	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketType: ticketType,
		Price:      price,
	}
	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}

// Validate checks if the ticket data is valid.
func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return errors.New("ticket price must be greater than zero")
	}
	return nil
}
