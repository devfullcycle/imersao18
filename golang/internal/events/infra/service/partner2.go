package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Partner2 struct {
	BaseURL string
}

type Partner2ReservationRequest struct {
	Lugares      []string `json:"lugares"`
	TipoIngresso string   `json:"tipo_ingresso"`
	Email        string   `json:"email"`
}

type Partner2ReservationResponse struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	Lugar        string `json:"lugar"`
	TipoIngresso string `json:"tipo_ingresso"`
	Status       string `json:"estado"`
	EventID      string `json:"evento_id"`
}

func (p *Partner2) MakeReservation(req *ReservationRequest) ([]ReservationResponse, error) {
	TipoIngresso := req.TicketKind
	if TipoIngresso == "full" {
		TipoIngresso = "inteira"
	} else {
		TipoIngresso = "meia"
	}
	partnerReq := Partner2ReservationRequest{
		Lugares:      req.Spots,
		TipoIngresso: TipoIngresso,
		Email:        req.Email,
	}

	body, err := json.Marshal(partnerReq)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/eventos/%s/reservar", p.BaseURL, req.EventID)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("reservation failed with status code: %d", httpResp.StatusCode)
	}

	var partnerResp []Partner2ReservationResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&partnerResp); err != nil {
		return nil, err
	}

	responses := make([]ReservationResponse, len(partnerResp))
	for i, r := range partnerResp {
		responses[i] = ReservationResponse{
			ID:     r.ID,
			Spot:   r.Lugar,
			Status: r.Status,
		}
	}

	return responses, nil
}
