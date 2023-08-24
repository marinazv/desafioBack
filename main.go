package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/marinazv/desafioBack/internal"
)

const (
	filename = "./tickets.csv"
)

func main() {
	// Recuperamos los errores para no romper el programa
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	storageTickets := internal.Storage{
		Tickets: ReadFile(filename),
	}

	fmt.Println("Ingrese un pais de destino")
	scanner := bufio.NewScanner(os.Stdin)

	// Escanear la próxima línea de texto
	scanner.Scan()

	// Obtener el texto ingresado por el usuario
	pais := scanner.Text()

	totalDestino, _ := storageTickets.GetTotalTickets(pais)
	fmt.Printf("Para el pais  %s  el total de ticketes fue %v \n", pais, totalDestino)

	fmt.Println("Ingrese el periodo (madrugada,mañana, tarde o noche) para el cual desea conocer la cantidad de pasajeros")

	scanner.Scan()
	periodo := scanner.Text()
	totalPeriodo, _ := storageTickets.GetCountByPeriod(periodo)

	fmt.Printf("Para el periodo  %s  el total de viajeros  fue %v", periodo, totalPeriodo)

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
			//horaVuelo, _ := time.Parse("2006-01-02 15:04:05", file[4]) // Ajustar el formato según cómo se almacene la hora en el archivo
			ticket := internal.Ticket{
				ID:          file[0],
				Nombre:      file[1],
				Email:       file[2],
				PaisDestino: file[3],
				HoraVuelo:   file[4],
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
