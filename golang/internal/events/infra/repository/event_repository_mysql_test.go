package repository

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/devfullcycle/imersao18/golang/internal/events/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("mysql", "test_user:test_password@tcp(localhost:3306)/")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Run tests
	m.Run()
}

func setupTestDB() error {
	_, err := db.Exec("DROP DATABASE IF EXISTS test_db")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE DATABASE test_db")
	if err != nil {
		return err
	}

	_, err = db.Exec("USE test_db")
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE events (
			id VARCHAR(36) NOT NULL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			location VARCHAR(255) NOT NULL,
			organization VARCHAR(255) NOT NULL,
			rating VARCHAR(10) NOT NULL,
			date DATETIME NOT NULL,
			image_url VARCHAR(255) NOT NULL,
			capacity INT NOT NULL,
			price FLOAT NOT NULL,
			partner_id INT NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE spots (
			id VARCHAR(36) NOT NULL PRIMARY KEY,
			event_id VARCHAR(36) NOT NULL,
			name VARCHAR(10) NOT NULL,
			status VARCHAR(10) NOT NULL,
			ticket_id VARCHAR(36),
			FOREIGN KEY (event_id) REFERENCES events(id)
		)
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE tickets (
			id VARCHAR(36) NOT NULL PRIMARY KEY,
			event_id VARCHAR(36) NOT NULL,
			spot_id VARCHAR(36) NOT NULL,
			ticket_kind VARCHAR(10) NOT NULL,
			price FLOAT NOT NULL,
			FOREIGN KEY (event_id) REFERENCES events(id),
			FOREIGN KEY (spot_id) REFERENCES spots(id)
		)
	`)
	return err
}

func TestMysqlEventRepository(t *testing.T) {
	repo := &mysqlEventRepository{db: db}

	t.Run("CreateEvent", func(t *testing.T) {
		err := setupTestDB()
		assert.Nil(t, err)

		eventID := uuid.New().String()
		event := &domain.Event{
			ID:           eventID,
			Name:         "Concert",
			Location:     "Stadium",
			Organization: "Music Inc.",
			Rating:       domain.RatingLivre,
			Date:         time.Now().Add(24 * time.Hour),
			ImageURL:     "http://example.com/image.jpg",
			Capacity:     100,
			Price:        50.0,
			PartnerID:    1,
		}
		err = repo.CreateEvent(event)
		assert.Nil(t, err)

		// Verify the event was created
		storedEvent, err := repo.FindEventByID(eventID)
		assert.Nil(t, err)
		assert.Equal(t, event.ID, storedEvent.ID)
		assert.Equal(t, event.Name, storedEvent.Name)
	})

	t.Run("CreateSpot", func(t *testing.T) {
		err := setupTestDB()
		assert.Nil(t, err)

		// Create an event to associate the spot with
		eventID := uuid.New().String()
		event := &domain.Event{
			ID:           eventID,
			Name:         "Concert",
			Location:     "Stadium",
			Organization: "Music Inc.",
			Rating:       domain.RatingLivre,
			Date:         time.Now().Add(24 * time.Hour),
			ImageURL:     "http://example.com/image.jpg",
			Capacity:     100,
			Price:        50.0,
			PartnerID:    1,
		}
		err = repo.CreateEvent(event)
		assert.Nil(t, err)

		spotID := uuid.New().String()
		spot := &domain.Spot{
			ID:      spotID,
			EventID: eventID,
			Name:    "A1",
			Status:  domain.SpotStatusAvailable,
		}
		err = repo.CreateSpot(spot)
		assert.Nil(t, err)

		// Verify the spot was created
		storedSpot, err := repo.FindSpotByID(spotID)
		assert.Nil(t, err)
		assert.Equal(t, spot.ID, storedSpot.ID)
		assert.Equal(t, spot.Name, storedSpot.Name)
	})

	t.Run("CreateTicket", func(t *testing.T) {
		err := setupTestDB()
		assert.Nil(t, err)

		// Create an event and a spot to associate the ticket with
		eventID := uuid.New().String()
		event := &domain.Event{
			ID:           eventID,
			Name:         "Concert",
			Location:     "Stadium",
			Organization: "Music Inc.",
			Rating:       domain.RatingLivre,
			Date:         time.Now().Add(24 * time.Hour),
			ImageURL:     "http://example.com/image.jpg",
			Capacity:     100,
			Price:        50.0,
			PartnerID:    1,
		}
		err = repo.CreateEvent(event)
		assert.Nil(t, err)

		spotID := uuid.New().String()
		spot := &domain.Spot{
			ID:      spotID,
			EventID: eventID,
			Name:    "A1",
			Status:  domain.SpotStatusAvailable,
		}
		err = repo.CreateSpot(spot)
		assert.Nil(t, err)

		ticketID := uuid.New().String()
		ticket := &domain.Ticket{
			ID:         ticketID,
			EventID:    eventID,
			Spot:       spot,
			TicketKind: domain.TicketKindFull,
			Price:      50.0,
		}
		err = repo.CreateTicket(ticket)
		assert.Nil(t, err)

		// Verify the ticket was created
		storedSpot, err := repo.FindSpotByID(spotID)
		assert.Nil(t, err)
		assert.Equal(t, ticket.ID, storedSpot.TicketID)
	})
}
