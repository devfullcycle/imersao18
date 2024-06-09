package usecase

import (
	"strconv"

	"github.com/devfullcycle/imersao18/golang/internal/events/domain"
	"github.com/devfullcycle/imersao18/golang/internal/events/infra/service"
)

type BuyTicketsInputDTO struct {
	EventID    string   `json:"event_id"`
	Spots      []string `json:"spots"`
	TicketType string   `json:"ticket_type"`
	CardHash   string   `json:"card_hash"`
	Email      string   `json:"email"`
}

type BuyTicketsOutputDTO struct {
	Tickets []TicketDTO `json:"tickets"`
}

type BuyTicketsUseCase struct {
	repo           domain.EventRepository
	partnerFactory service.PartnerFactory
}

func NewBuyTicketsUseCase(repo domain.EventRepository, partnerFactory service.PartnerFactory) *BuyTicketsUseCase {
	return &BuyTicketsUseCase{
		repo:           repo,
		partnerFactory: partnerFactory,
	}
}

func (uc *BuyTicketsUseCase) Execute(input BuyTicketsInputDTO) (*BuyTicketsOutputDTO, error) {
	// Verify the event
	event, err := uc.repo.FindEventByID(input.EventID)
	if err != nil {
		return nil, err
	}

	// Create reservation request
	req := &service.ReservationRequest{
		EventID:    input.EventID,
		Spots:      input.Spots,
		TicketType: input.TicketType,
		CardHash:   input.CardHash,
		Email:      input.Email,
	}

	// Get the partner service
	partnerService, err := uc.partnerFactory.CreatePartner(event.PartnerID)
	if err != nil {
		return nil, err
	}

	// Reserve spots using the partner service
	reservationResponse, err := partnerService.MakeReservation(req)
	if err != nil {
		return nil, err
	}

	// Save tickets in the database
	tickets := make([]domain.Ticket, len(reservationResponse))
	for i, reservation := range reservationResponse {
		spot, err := uc.repo.FindSpotByID(reservation.Spot)
		if err != nil {
			return nil, err
		}

		ticket := &domain.Ticket{
			ID:         strconv.Itoa(reservation.ID),
			EventID:    input.EventID,
			Spot:       spot,
			TicketType: domain.TicketType(input.TicketType),
			Price:      event.Price,
		}

		if input.TicketType == string(domain.TicketTypeHalf) {
			ticket.Price /= 2
		}

		err = uc.repo.CreateTicket(ticket)
		if err != nil {
			return nil, err
		}

		spot.Reserve(ticket.ID)
		err = uc.repo.ReserveSpot(spot.ID, ticket.ID)
		if err != nil {
			return nil, err
		}

		tickets[i] = *ticket
	}

	ticketDTOs := make([]TicketDTO, len(tickets))
	for i, ticket := range tickets {
		ticketDTOs[i] = TicketDTO{
			ID:         ticket.ID,
			SpotID:     ticket.Spot.ID,
			TicketType: string(ticket.TicketType),
			Price:      ticket.Price,
		}
	}

	return &BuyTicketsOutputDTO{Tickets: ticketDTOs}, nil
}
