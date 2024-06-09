package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartnerFactory_CreatePartner(t *testing.T) {
	partnerBaseURLs := map[int]string{
		1: "http://localhost:9000/api1",
		2: "http://localhost:9000/api2",
	}
	factory := NewPartnerFactory(partnerBaseURLs)

	partner1, err := factory.CreatePartner(1)
	assert.NoError(t, err)
	assert.IsType(t, &Partner1{}, partner1)
	assert.Equal(t, "http://localhost:9000/api1", partner1.(*Partner1).BaseURL)

	partner2, err := factory.CreatePartner(2)
	assert.NoError(t, err)
	assert.IsType(t, &Partner2{}, partner2)
	assert.Equal(t, "http://localhost:9000/api2", partner2.(*Partner2).BaseURL)

	_, err = factory.CreatePartner(3)
	assert.Error(t, err)
	assert.Equal(t, "unsupported partner ID: 3", err.Error())
}
