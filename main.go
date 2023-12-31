package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/marinazv/desafioBack/internal"
)

const (
	filename = "./tickets.csv"
)

func main() {
	// Crear canales para comunicarnos con las goroutines
	canalGetTotalTickets := make(chan int)
	canalGetCountByPeriod := make(chan int)
	canalPorcentaje := make(chan float64)

	canalErr := make(chan error)
	// creo la variable para las esperas de los go rutines
	var wg sync.WaitGroup
	wg.Add(3)

	// Recuperamos los errores para no romper el programa
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		defer close(canalErr)
		for err := range canalErr {
			log.Println("Error:", err)
		}
	}()

	//defino para tener la base de datos en memoria:
	storageTickets := internal.Storage{
		Tickets: ReadFile(filename),
	}

	fmt.Println("Ingrese un país de destino")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	destination := scanner.Text()

	fmt.Println("Ingrese el periodo (madrugada, mañana, tarde o noche) para el cual desea conocer la cantidad de pasajeros")
	scanner.Scan()
	period := scanner.Text()

	// Goroutine para obtener el total de tickets por destino
	go func() {
		defer wg.Done()
		fmt.Println("Procesando gorutine 1")
		// obtener el total de tickets por destino/
		total, err := storageTickets.GetTotalTickets(destination)
		if err != nil {
			canalErr <- err
			return
		}

		canalGetTotalTickets <- total
		fmt.Println(" Terminando de Procesar gorutine 1")
	}()

	// Goroutine para obtener la cantidad de pasajeros por periodo
	go func() {
		defer wg.Done()
		fmt.Println("Procesando gorutine 2")
		// obtener el total de tickets por destiperiodo
		total, err := storageTickets.GetCountByPeriod(period)
		if err != nil {
			canalErr <- err
			return
		}
		canalGetCountByPeriod <- total
		fmt.Println(" Terminando de Procesar gorutine 2")
	}()

	// Goroutine para obtener porcentaje
	go func() {
		defer wg.Done()
		fmt.Println("Procesando gorutine 3")
		// obtener el porcentaje
		porcentaje, err := storageTickets.AverageDestination(destination, &storageTickets.Tickets)
		if err != nil {
			canalErr <- err
			return
		}
		canalPorcentaje <- porcentaje
		fmt.Println(" Terminando de Procesar gorutine 3")
	}()

	// asigno a variables lo que traen los canales
	GetTotalTickets := <-canalGetTotalTickets
	fmt.Printf("la cantidad de tickets para el destino %v es %v\n", destination, GetTotalTickets)
	CountByPeriod := <-canalGetCountByPeriod
	fmt.Printf("La canitdad de tickets para el periodo %v  es %v\n", period, CountByPeriod)
	AverageDestination := <-canalPorcentaje
	fmt.Printf("El porcentaje por destino es: %v\n", AverageDestination)
	// espero a que todas la go rutines terminen
	wg.Wait()

	// Cerrar los canales adecuadamente
	close(canalGetTotalTickets)
	close(canalGetCountByPeriod)
	close(canalPorcentaje)

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
