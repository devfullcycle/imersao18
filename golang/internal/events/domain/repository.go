package domain

// EventRepository defines the interface for interacting with event and spot data.
type EventRepository interface {
	ListEvents() ([]Event, error)
	FindEventByID(eventID string) (*Event, error)
	CreateEvent(event *Event) error
	FindSpotByID(spotID string) (*Spot, error)
	CreateTicket(ticket *Ticket) error
	ReserveSpot(spotID, ticketID string) error
}
