package usecase

import (
	"github.com/devfullcycle/imersao18/golang/internal/events/domain"
)

type GetEventInputDTO struct {
	ID string
}

type GetEventOutputDTO struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Location     string      `json:"location"`
	Organization string      `json:"organization"`
	Rating       string      `json:"rating"`
	Date         string      `json:"date"`
	Capacity     int         `json:"capacity"`
	Price        float64     `json:"price"`
	PartnerID    int         `json:"partner_id"`
	Spots        []SpotDTO   `json:"spots"`
	Tickets      []TicketDTO `json:"tickets"`
}

type SpotDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	TicketID string `json:"ticket_id"`
}

type TicketDTO struct {
	ID         string  `json:"id"`
	SpotID     string  `json:"spot_id"`
	TicketType string  `json:"ticket_type"`
	Price      float64 `json:"price"`
}

type GetEventUseCase struct {
	repo domain.EventRepository
}

func NewGetEventUseCase(repo domain.EventRepository) *GetEventUseCase {
	return &GetEventUseCase{repo: repo}
}

func (uc *GetEventUseCase) Execute(input GetEventInputDTO) (*GetEventOutputDTO, error) {
	event, err := uc.repo.FindEventByID(input.ID)
	if err != nil {
		return nil, err
	}

	spots := make([]SpotDTO, len(event.Spots))
	for i, spot := range event.Spots {
		spots[i] = SpotDTO{
			ID:       spot.ID,
			Name:     spot.Name,
			Status:   string(spot.Status),
			TicketID: spot.TicketID,
		}
	}

	tickets := make([]TicketDTO, len(event.Tickets))
	for i, ticket := range event.Tickets {
		tickets[i] = TicketDTO{
			ID:         ticket.ID,
			SpotID:     ticket.Spot.ID,
			TicketType: string(ticket.TicketType),
			Price:      ticket.Price,
		}
	}

	return &GetEventOutputDTO{
		ID:           event.ID,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date.Format("2006-01-02 15:04:05"),
		Capacity:     event.Capacity,
		Price:        event.Price,
		PartnerID:    event.PartnerID,
		Spots:        spots,
		Tickets:      tickets,
	}, nil
}
