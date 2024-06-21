package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPartner2_MakeReservation_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api2/eventos/1/reservar", r.URL.Path)
		w.WriteHeader(http.StatusCreated) // Mudança aqui para http.StatusCreated
		w.Write([]byte(`[{"id": "1", "email": "user2@test.com", "lugar": "A1", "tipo_ingresso": "inteira", "estado": "reservado", "evento_id": "1"}]`))
	}))
	defer server.Close()

	partner := &Partner2{BaseURL: server.URL + "/api2"}
	req := &ReservationRequest{
		EventID:    "1",
		Spots:      []string{"A1"},
		TicketKind: "inteira",
		Email:      "user2@test.com",
	}

	resp, err := partner.MakeReservation(req)
	require.NoError(t, err)
	require.Len(t, resp, 1)
	assert.Equal(t, "1", resp[0].ID) // Mudança aqui para comparar strings
	assert.Equal(t, "A1", resp[0].Spot)
	assert.Equal(t, "reservado", resp[0].Status)
}

func TestPartner2_MakeReservation_Failure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api2/eventos/1/reservar", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest) // Mudança aqui para simular falha
		w.Write([]byte(`{"error": "invalid request"}`))
	}))
	defer server.Close()

	partner := &Partner2{BaseURL: server.URL + "/api2"}
	req := &ReservationRequest{
		EventID:    "1",
		Spots:      []string{"A1"},
		TicketKind: "inteira",
		Email:      "user2@test.com",
	}

	resp, err := partner.MakeReservation(req)
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "reservation failed with status code: 400")
}
