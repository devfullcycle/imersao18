package usecase

import (
	"testing"
	"time"

	"github.com/devfullcycle/imersao18/golang/internal/events/domain"
	"github.com/devfullcycle/imersao18/golang/internal/events/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetEventUseCase(t *testing.T) {
	mockRepo := new(repository.MockEventRepository)
	getEventUseCase := NewGetEventUseCase(mockRepo)

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
		},
		Tickets: []domain.Ticket{
			{ID: "1", EventID: eventID, Spot: &domain.Spot{ID: "1"}, TicketType: domain.TicketTypeFull, Price: 50.0},
		},
	}

	// Mock the repository to expect the call to FindEventByID
	mockRepo.On("FindEventByID", eventID).Return(mockEvent, nil)

	// Define the input DTO
	input := GetEventInputDTO{
		ID: eventID,
	}

	// Execute the use case
	output, err := getEventUseCase.Execute(input)

	// Assertions
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockEvent.ID, output.ID)
	assert.Equal(t, mockEvent.Name, output.Name)
	assert.Equal(t, mockEvent.Location, output.Location)
	assert.Equal(t, mockEvent.Organization, output.Organization)
	assert.Equal(t, string(mockEvent.Rating), output.Rating)
	assert.Equal(t, mockEvent.Date.Format("2006-01-02 15:04:05"), output.Date)
	assert.Equal(t, mockEvent.Capacity, output.Capacity)
	assert.Equal(t, mockEvent.Price, output.Price)
	assert.Equal(t, mockEvent.PartnerID, output.PartnerID)

	assert.Equal(t, len(mockEvent.Spots), len(output.Spots))
	assert.Equal(t, mockEvent.Spots[0].ID, output.Spots[0].ID)
	assert.Equal(t, mockEvent.Spots[0].Name, output.Spots[0].Name)
	assert.Equal(t, string(mockEvent.Spots[0].Status), output.Spots[0].Status)
	assert.Equal(t, mockEvent.Spots[0].TicketID, output.Spots[0].TicketID)

	assert.Equal(t, len(mockEvent.Tickets), len(output.Tickets))
	assert.Equal(t, mockEvent.Tickets[0].ID, output.Tickets[0].ID)
	assert.Equal(t, mockEvent.Tickets[0].Spot.ID, output.Tickets[0].SpotID)
	assert.Equal(t, string(mockEvent.Tickets[0].TicketType), output.Tickets[0].TicketType)
	assert.Equal(t, mockEvent.Tickets[0].Price, output.Tickets[0].Price)

	// Assert that the mock repository was called
	mockRepo.AssertExpectations(t)
}
