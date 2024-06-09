package usecase

import (
	"testing"
	"time"

	"github.com/devfullcycle/imersao18/golang/internal/events/domain"
	"github.com/devfullcycle/imersao18/golang/internal/events/infra/repository"
	"github.com/devfullcycle/imersao18/golang/internal/events/infra/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPartnerService struct {
	mock.Mock
}

func (m *MockPartnerService) MakeReservation(req *service.ReservationRequest) ([]service.ReservationResponse, error) {
	args := m.Called(req)
	return args.Get(0).([]service.ReservationResponse), args.Error(1)
}

type MockPartnerFactory struct {
	mock.Mock
}

func (m *MockPartnerFactory) CreatePartner(partnerID int) (service.Partner, error) {
	args := m.Called(partnerID)
	return args.Get(0).(service.Partner), args.Error(1)
}

func TestBuyTicketsUseCase(t *testing.T) {
	mockRepo := new(repository.MockEventRepository)
	mockPartnerService := new(MockPartnerService)
	mockPartnerFactory := new(MockPartnerFactory)

	// Configurações da factory mock
	mockPartnerFactory.On("CreatePartner", 1).Return(mockPartnerService, nil)

	buyTicketsUseCase := NewBuyTicketsUseCase(mockRepo, mockPartnerFactory)

	eventID := "1"
	eventDate := time.Now().Add(24 * time.Hour)
	mockEvent := &domain.Event{
		ID:           eventID,
		Name:         "Concert",
		Location:     "Stadium",
		Organization: "Music Inc.",
		Rating:       domain.RatingLivre,
		Date:         eventDate,
		ImageURL:     "http://example.com/image.jpg",
		Capacity:     100,
		Price:        50.0,
		PartnerID:    1,
		Spots: []domain.Spot{
			{ID: "1", EventID: eventID, Name: "A1", Status: domain.SpotStatusAvailable},
			{ID: "2", EventID: eventID, Name: "A2", Status: domain.SpotStatusAvailable},
		},
		Tickets: []domain.Ticket{},
	}

	// Mock the repository to expect the call to FindEventByID
	mockRepo.On("FindEventByID", eventID).Return(mockEvent, nil)
	mockRepo.On("FindSpotByID", "1").Return(&mockEvent.Spots[0], nil)
	mockRepo.On("FindSpotByID", "2").Return(&mockEvent.Spots[1], nil)
	mockRepo.On("CreateTicket", mock.AnythingOfType("*domain.Ticket")).Return(nil)
	mockRepo.On("ReserveSpot", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	// Mock the partner service to expect the call to MakeReservation
	mockPartnerService.On("MakeReservation", mock.AnythingOfType("*service.ReservationRequest")).Return([]service.ReservationResponse{
		{ID: 1, Spot: "1"},
		{ID: 2, Spot: "2"},
	}, nil)

	// Define the input DTO
	input := BuyTicketsInputDTO{
		EventID:    eventID,
		Spots:      []string{"A1", "A2"},
		TicketType: "full",
		CardHash:   "hash_do_cartao",
		Email:      "test@test.com",
	}

	// Execute the use case
	output, err := buyTicketsUseCase.Execute(input)

	// Assertions
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, 2, len(output.Tickets))

	assert.Equal(t, "1", output.Tickets[0].ID)
	assert.Equal(t, "1", output.Tickets[0].SpotID)
	assert.Equal(t, "full", output.Tickets[0].TicketType)
	assert.Equal(t, 50.0, output.Tickets[0].Price)

	assert.Equal(t, "2", output.Tickets[1].ID)
	assert.Equal(t, "2", output.Tickets[1].SpotID)
	assert.Equal(t, "full", output.Tickets[1].TicketType)
	assert.Equal(t, 50.0, output.Tickets[1].Price)

	// Assert that the mock repository and partner service were called
	mockRepo.AssertExpectations(t)
	mockPartnerFactory.AssertExpectations(t)
	mockPartnerService.AssertExpectations(t)
}
