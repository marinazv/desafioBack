package internal

import (
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("Not found")
)

type Ticket struct {
	ID          string
	Nombre      string
	Email       string
	PaisDestino string
	HoraVuelo   time.Time
	Precio      string
}

type Storage struct {
	Tickets []Ticket
}

// GetTotalTickets devuelve la cantidad total de tickets para un destino específico.
func (s *Storage) GetTotalTickets(destination string) (int, error) {
	totalTickets := 0

	for _, ticket := range s.Tickets {
		if ticket.PaisDestino == destination {
			totalTickets++
		}
	}

	if totalTickets > 0 {
		return totalTickets, nil
	} else {
		return 0, ErrNotFound
	}
}
