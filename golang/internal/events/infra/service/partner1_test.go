package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPartner1_MakeReservation_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api1/events/1/reserve", r.URL.Path)
		w.WriteHeader(http.StatusCreated) // Mudança aqui para http.StatusCreated
		w.Write([]byte(`[{"id": "1", "email": "user1@test.com", "spot": "A1", "ticket_kind": "full", "status": "reserved", "event_id": "1"}]`))
	}))
	defer server.Close()

	partner := &Partner1{BaseURL: server.URL + "/api1"}
	req := &ReservationRequest{
		EventID:    "1",
		Spots:      []string{"A1"},
		TicketKind: "full",
		Email:      "user1@test.com",
	}

	resp, err := partner.MakeReservation(req)
	require.NoError(t, err)
	require.Len(t, resp, 1)
	assert.Equal(t, "1", resp[0].ID) // Mudança aqui para comparar strings
	assert.Equal(t, "A1", resp[0].Spot)
	assert.Equal(t, "reserved", resp[0].Status)
}

func TestPartner1_MakeReservation_Failure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api1/events/1/reserve", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest) // Mudança aqui para simular falha
		w.Write([]byte(`{"error": "invalid request"}`))
	}))
	defer server.Close()

	partner := &Partner1{BaseURL: server.URL + "/api1"}
	req := &ReservationRequest{
		EventID:    "1",
		Spots:      []string{"A1"},
		TicketKind: "full",
		Email:      "user1@test.com",
	}

	resp, err := partner.MakeReservation(req)
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "reservation failed with status code: 400")
}
