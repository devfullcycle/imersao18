package usecase

import (
	"testing"
	"time"

	"github.com/devfullcycle/imersao18/golang/internal/events/domain"
	"github.com/devfullcycle/imersao18/golang/internal/events/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestListEventsUseCase(t *testing.T) {
	mockRepo := new(repository.MockEventRepository)
	listEventsUseCase := NewListEventsUseCase(mockRepo)

	// Mock events data
	eventID := "1"
	eventDate := time.Now().Add(24 * time.Hour)
	mockEvents := []domain.Event{
		{
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
		},
	}

	// Mock the repository to expect the call to ListEvents
	mockRepo.On("ListEvents").Return(mockEvents, nil)

	// Execute the use case
	output, err := listEventsUseCase.Execute()

	// Assertions
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, len(mockEvents), len(output.Events))

	// Verify the first event
	assert.Equal(t, mockEvents[0].ID, output.Events[0].ID)
	assert.Equal(t, mockEvents[0].Name, output.Events[0].Name)
	assert.Equal(t, mockEvents[0].Location, output.Events[0].Location)
	assert.Equal(t, mockEvents[0].Organization, output.Events[0].Organization)
	assert.Equal(t, string(mockEvents[0].Rating), output.Events[0].Rating)
	assert.Equal(t, mockEvents[0].Date.Format("2006-01-02 15:04:05"), output.Events[0].Date)
	assert.Equal(t, mockEvents[0].Capacity, output.Events[0].Capacity)
	assert.Equal(t, mockEvents[0].Price, output.Events[0].Price)
	assert.Equal(t, mockEvents[0].PartnerID, output.Events[0].PartnerID)

	assert.Equal(t, len(mockEvents[0].Spots), len(output.Events[0].Spots))
	assert.Equal(t, mockEvents[0].Spots[0].ID, output.Events[0].Spots[0].ID)
	assert.Equal(t, mockEvents[0].Spots[0].Name, output.Events[0].Spots[0].Name)
	assert.Equal(t, string(mockEvents[0].Spots[0].Status), output.Events[0].Spots[0].Status)
	assert.Equal(t, mockEvents[0].Spots[0].TicketID, output.Events[0].Spots[0].TicketID)

	assert.Equal(t, len(mockEvents[0].Tickets), len(output.Events[0].Tickets))
	assert.Equal(t, mockEvents[0].Tickets[0].ID, output.Events[0].Tickets[0].ID)
	assert.Equal(t, mockEvents[0].Tickets[0].Spot.ID, output.Events[0].Tickets[0].SpotID)
	assert.Equal(t, string(mockEvents[0].Tickets[0].TicketType), output.Events[0].Tickets[0].TicketType)
	assert.Equal(t, mockEvents[0].Tickets[0].Price, output.Events[0].Tickets[0].Price)

	// Assert that the mock repository was called
	mockRepo.AssertExpectations(t)
}
