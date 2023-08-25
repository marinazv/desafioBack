package internal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrNotFoundTicket = errors.New("Not found ticket")
	ErrNotFoundTime   = errors.New("Not found Time")
)

type Ticket struct {
	ID          string
	Nombre      string
	Email       string
	PaisDestino string
	HoraVuelo   string
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
		return 0, ErrNotFoundTicket
	}
}

// Calcula cuantas personas viajan en madrugada tarde y noche
func (s *Storage) GetCountByPeriod(time string) (int, error) {

	totalPersonas := 0

	for i := 1; i < len(s.Tickets); i++ {
		ticket := s.Tickets[i]
		horaString := strings.Split(string(ticket.HoraVuelo), ":")[0]
		hora, err := strconv.Atoi(horaString)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}

		switch {
		case time == "madrugada":
			if hora >= 0 && hora < 7 {
				totalPersonas++
			}

		case time == "mañana":
			if hora >= 7 && hora < 13 {
				totalPersonas++
			}

		case time == "tarde":
			if hora >= 13 && hora < 20 {
				totalPersonas++
			}

		case time == "noche":
			if hora >= 20 && hora <= 24 {
				totalPersonas++
			}

		default:
			totalPersonas = 0
		}
	}

	if totalPersonas > 0 {
		return totalPersonas, nil
	} else {
		return 0, ErrNotFoundTime
	}

}
