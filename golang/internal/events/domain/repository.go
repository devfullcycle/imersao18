package domain

type EventRepository interface {
	ListEvents() ([]Event, error)
	FindEventByID(eventID string) (*Event, error)
	FindSpotsByEventID(eventID string) ([]*Spot, error)
	FindSpotByName(eventID, spotName string) (*Spot, error) // Atualizado
	CreateEvent(event *Event) error
	CreateSpot(spot *Spot) error
	CreateTicket(ticket *Ticket) error
	ReserveSpot(spotID, ticketID string) error
}
