package usecase

import (
	"testing"
	"time"

	"github.com/devfullcycle/imersao18/golang/internal/events/domain"
	"github.com/devfullcycle/imersao18/golang/internal/events/infra/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateEventUseCase(t *testing.T) {
	mockRepo := new(repository.MockEventRepository)
	createEventUseCase := NewCreateEventUseCase(mockRepo)

	// Define the input DTO
	input := CreateEventInputDTO{
		Name:         "Concert",
		Location:     "Stadium",
		Organization: "Music Inc.",
		Rating:       string(domain.RatingLivre),
		Date:         time.Now().Add(24 * time.Hour),
		Capacity:     100,
		Price:        50.0,
		PartnerID:    1,
	}

	// Mock the repository to expect the call to CreateEvent
	mockRepo.On("CreateEvent", mock.AnythingOfType("*domain.Event")).Return(nil)

	// Execute the use case
	output, err := createEventUseCase.Execute(input)

	// Assertions
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Location, output.Location)
	assert.Equal(t, input.Organization, output.Organization)
	assert.Equal(t, input.Rating, output.Rating)
	assert.WithinDuration(t, input.Date, output.Date, time.Second)
	assert.Equal(t, input.Capacity, output.Capacity)
	assert.Equal(t, input.Price, output.Price)
	assert.Equal(t, input.PartnerID, output.PartnerID)

	// Assert that the mock repository was called
	mockRepo.AssertExpectations(t)
}
