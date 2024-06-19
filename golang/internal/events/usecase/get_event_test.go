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

	// Assert that the mock repository was called
	mockRepo.AssertExpectations(t)
}
