package usecase

import (
	"time"

	"github.com/devfullcycle/imersao18/golang/internal/events/domain"
)

// CreateEventInputDTO represents the input data required to create an event.
type CreateEventInputDTO struct {
	Name         string    `json:"name"`
	Location     string    `json:"location"`
	Organization string    `json:"organization"`
	Rating       string    `json:"rating"`
	Date         time.Time `json:"date"`
	Capacity     int       `json:"capacity"`
	ImageURL     string    `json:"image_url"`
	Price        float64   `json:"price"`
	PartnerID    int       `json:"partner_id"`
}

// CreateEventOutputDTO represents the output data after creating an event.
type CreateEventOutputDTO struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Location     string    `json:"location"`
	Organization string    `json:"organization"`
	Rating       string    `json:"rating"`
	Date         time.Time `json:"date"`
	Capacity     int       `json:"capacity"`
	ImageURL     string    `json:"image_url"`
	Price        float64   `json:"price"`
	PartnerID    int       `json:"partner_id"`
}

// CreateEventUseCase defines the use case for creating events.
type CreateEventUseCase struct {
	repo domain.EventRepository
}

// NewCreateEventUseCase creates a new instance of CreateEventUseCase.
func NewCreateEventUseCase(repo domain.EventRepository) *CreateEventUseCase {
	return &CreateEventUseCase{repo: repo}
}

// Execute executes the use case to create an event.
func (uc *CreateEventUseCase) Execute(input CreateEventInputDTO) (CreateEventOutputDTO, error) {
	event, err := domain.NewEvent(
		input.Name,
		input.Location,
		input.Organization,
		domain.Rating(input.Rating),
		input.Date,
		input.Capacity,
		input.Price,
		input.ImageURL,
		input.PartnerID,
	)
	if err != nil {
		return CreateEventOutputDTO{}, err
	}

	err = uc.repo.CreateEvent(event)
	if err != nil {
		return CreateEventOutputDTO{}, err
	}

	output := CreateEventOutputDTO{
		ID:           event.ID,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date,
		Capacity:     event.Capacity,
		ImageURL:     event.ImageURL,
		Price:        event.Price,
		PartnerID:    event.PartnerID,
	}

	return output, nil
}
