package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/marinazv/desafioBack/internal"
)

const (
	filename = "./tickets.csv"
)

func main() {

	ticket := internal.Ticket{
		ID:          "T123",
		Nombre:      "Ejemplo Cliente",
		Email:       "cliente@example.com",
		PaisDestino: "Destinolandia",
		HoraVuelo:   time.Now(), // Usar la hora actual
		Precio:      "$500",
	}

	// Imprimir el ticket
	fmt.Printf("ID: %s\n", ticket.ID)
	fmt.Printf("Nombre: %s\n", ticket.Nombre)
	fmt.Printf("Email: %s\n", ticket.Email)
	fmt.Printf("País Destino: %s\n", ticket.PaisDestino)
	fmt.Printf("Hora de Vuelo: %s\n", ticket.HoraVuelo.Format("2006-01-02 15:04:05"))
	fmt.Printf("Precio: %s\n", ticket.Precio)

	tickets := ReadFile(filename)

	for _, ticket := range tickets {
		fmt.Printf("ID: %s\n", ticket.ID)
		fmt.Printf("Nombre: %s\n", ticket.Nombre)
		fmt.Printf("Email: %s\n", ticket.Email)
		fmt.Printf("PaisDestino: %s\n", ticket.PaisDestino)
		fmt.Printf("HoraVuelo: %s\n", ticket.HoraVuelo.Format(time.RFC3339))
		fmt.Printf("Precio: %s\n", ticket.Precio)
		fmt.Println("------------")
	}

	HolaMundo1()

}

func ReadFile(filename string) []internal.Ticket {
	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	data := strings.Split(string(file), "\n")

	var resultado []internal.Ticket
	for i := 0; i < len(data); i++ {

		if len(data[i]) > 0 {
			file := strings.Split(string(data[i]), ",")
			horaVuelo, _ := time.Parse("2006-01-02 15:04:05", file[4]) // Ajustar el formato según cómo se almacene la hora en el archivo
			ticket := internal.Ticket{
				ID:          file[0],
				Nombre:      file[1],
				Email:       file[2],
				PaisDestino: file[3],
				HoraVuelo:   horaVuelo,
				Precio:      file[5],
			}
			resultado = append(resultado, ticket)
		}

	}

	return resultado

}

func HolaMundo1() {
	fmt.Println("hola mundo")
}
