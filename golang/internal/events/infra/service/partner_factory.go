package service

import "fmt"

type PartnerFactory interface {
	CreatePartner(partnerID int) (Partner, error)
}

type DefaultPartnerFactory struct {
	partnerBaseURLs map[int]string
}

func NewPartnerFactory(partnerBaseURLs map[int]string) PartnerFactory {
	return &DefaultPartnerFactory{partnerBaseURLs: partnerBaseURLs}
}

func (f *DefaultPartnerFactory) CreatePartner(partnerID int) (Partner, error) {
	baseURL, ok := f.partnerBaseURLs[partnerID]
	if !ok {
		return nil, fmt.Errorf("unsupported partner ID: %d", partnerID)
	}

	switch partnerID {
	case 1:
		return &Partner1{BaseURL: baseURL}, nil
	case 2:
		return &Partner2{BaseURL: baseURL}, nil
	default:
		return nil, fmt.Errorf("unsupported partner ID: %d", partnerID)
	}
}
