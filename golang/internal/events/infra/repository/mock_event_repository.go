package repository

import (
	"github.com/devfullcycle/imersao18/golang/internal/events/domain"
	"github.com/stretchr/testify/mock"
)

type MockEventRepository struct {
	mock.Mock
}

func (m *MockEventRepository) ListEvents() ([]domain.Event, error) {
	args := m.Called()
	return args.Get(0).([]domain.Event), args.Error(1)
}

func (m *MockEventRepository) FindEventByID(eventID string) (*domain.Event, error) {
	args := m.Called(eventID)
	return args.Get(0).(*domain.Event), args.Error(1)
}

func (m *MockEventRepository) FindSpotsByEventID(eventID string) ([]*domain.Spot, error) {
	args := m.Called(eventID)
	return args.Get(0).([]*domain.Spot), args.Error(1)
}

func (m *MockEventRepository) FindSpotByID(spotID string) (*domain.Spot, error) {
	args := m.Called(spotID)
	return args.Get(0).(*domain.Spot), args.Error(1)
}

func (m *MockEventRepository) FindSpotByName(eventID, name string) (*domain.Spot, error) {
	args := m.Called(eventID, name)
	return args.Get(0).(*domain.Spot), args.Error(1)
}

func (m *MockEventRepository) CreateEvent(event *domain.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockEventRepository) CreateSpot(spot *domain.Spot) error {
	args := m.Called(spot)
	return args.Error(0)
}

func (m *MockEventRepository) CreateTicket(ticket *domain.Ticket) error {
	args := m.Called(ticket)
	return args.Error(0)
}

func (m *MockEventRepository) ReserveSpot(spotID, ticketID string) error {
	args := m.Called(spotID, ticketID)
	return args.Error(0)
}
