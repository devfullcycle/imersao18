package service

import "github.com/stretchr/testify/mock"

type MockPartnerFactory struct {
	mock.Mock
}

func (m *MockPartnerFactory) CreatePartner(partnerID int) (Partner, error) {
	args := m.Called(partnerID)
	return args.Get(0).(Partner), args.Error(1)
}
