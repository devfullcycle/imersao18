package http

import (
	"encoding/json"
	"net/http"

	"github.com/devfullcycle/imersao18/golang/internal/events/usecase"
)

type EventsHandler struct {
	listEventsUseCase  *usecase.ListEventsUseCase
	getEventUseCase    *usecase.GetEventUseCase
	createEventUseCase *usecase.CreateEventUseCase
	buyTicketsUseCase  *usecase.BuyTicketsUseCase
	listSpotsUseCase   *usecase.ListSpotsUseCase
}

func NewEventsHandler(
	listEventsUseCase *usecase.ListEventsUseCase,
	getEventUseCase *usecase.GetEventUseCase,
	createEventUseCase *usecase.CreateEventUseCase,
	buyTicketsUseCase *usecase.BuyTicketsUseCase,
	listSpotsUseCase *usecase.ListSpotsUseCase,
) *EventsHandler {
	return &EventsHandler{
		listEventsUseCase:  listEventsUseCase,
		getEventUseCase:    getEventUseCase,
		createEventUseCase: createEventUseCase,
		buyTicketsUseCase:  buyTicketsUseCase,
		listSpotsUseCase:   listSpotsUseCase,
	}
}

// ListEvents godoc
// @Summary List all events
// @Description Get all events
// @Tags events
// @Produce json
// @Success 200 {object} usecase.ListEventsOutputDTO
// @Router /events [get]
func (h *EventsHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	output, err := h.listEventsUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// GetEvent godoc
// @Summary Get event by ID
// @Description Get a single event by ID
// @Tags events
// @Produce json
// @Param eventID path string true "Event ID"
// @Success 200 {object} usecase.GetEventOutputDTO
// @Router /events/{eventID} [get]
func (h *EventsHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")

	input := usecase.GetEventInputDTO{ID: eventID}
	output, err := h.getEventUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// CreateEvent godoc
// @Summary Create a new event
// @Description Create a new event
// @Tags events
// @Accept json
// @Produce json
// @Param event body usecase.CreateEventInputDTO true "Event data"
// @Success 201 {object} usecase.CreateEventOutputDTO
// @Router /events [post]
func (h *EventsHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateEventInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.createEventUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// BuyTickets godoc
// @Summary Buy tickets for an event
// @Description Buy tickets
// @Tags tickets
// @Accept json
// @Produce json
// @Param request body usecase.BuyTicketsInputDTO true "Buy tickets data"
// @Success 200 {object} usecase.BuyTicketsOutputDTO
// @Router /checkout [post]
func (h *EventsHandler) BuyTickets(w http.ResponseWriter, r *http.Request) {
	var input usecase.BuyTicketsInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.buyTicketsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// ListSpots godoc
// @Summary List spots for an event
// @Description Get all spots for an event
// @Tags spots
// @Produce json
// @Param eventID path string true "Event ID"
// @Success 200 {object} usecase.ListSpotsOutputDTO
// @Router /events/{eventID}/spots [get]
func (h *EventsHandler) ListSpots(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	input := usecase.ListSpotsInputDTO{EventID: eventID}
	output, err := h.listSpotsUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
