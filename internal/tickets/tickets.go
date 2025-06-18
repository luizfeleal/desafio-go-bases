package tickets

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/luizfeleal/desafio-go-bases/internal/tickets/domain"
)

var (
	fileName = "./tickets.csv"
)

type Ticket interface {
	GetTotalTickets(destination string) (int, error)
	GetCountryByPeriod(country string) (map[string]int, error)
	AverageDestination(destination string, total int) (int, error)
}
type repository struct {
	filePath string
	data     []domain.Ticket
}

type ticketsOpt func(*repository)

func WithFilePath(filePath string) ticketsOpt {
	return func(r *repository) {
		r.filePath = filePath
	}
}

func NewRepository(opt ...ticketsOpt) *repository {
	defaultRepository := repository{
		filePath: fileName,
	}
	for _, o := range opt {
		o(&defaultRepository)
	}

	file, err := os.ReadFile(defaultRepository.filePath)
	if err != nil {
		fmt.Println("Erro:", err)
		return nil
	}
	buf := bytes.NewBuffer(file)
	reader := csv.NewReader(buf)

	var basicData []domain.Ticket
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Erro ao ler linha:", err)
			break
		}
		var ticket domain.Ticket

		if err = ticket.NormalizedTicket(record); err != nil {
			return nil
		}
		basicData = append(basicData, ticket)

	}

	defaultRepository.data = basicData
	return &defaultRepository
}

func (t *repository) GetTotalTickets(destination string) (int, error) {
	totalTickets := t.data
	var ticketsFound int
	for _, ticket := range totalTickets {
		if ticket.Destination == destination {
			ticketsFound++
		}
	}

	return ticketsFound, nil
}

func (t *repository) GetCountryByPeriod(country string) (map[string]int, error) {

	ticketsCountByTime := map[string]int{
		"inicio_manha": 0,
		"manha":        0,
		"tarde":        0,
		"noite":        0,
	}
	for _, ticket := range t.data {
		if ticket.Destination == country {
			hourString := ticket.Arrival
			hourParts := strings.Split(hourString, ":")
			hour, err := strconv.Atoi(hourParts[0])
			if err != nil {
				fmt.Errorf("A hora precisa ser um valor válido")
			}
			switch {
			case hour >= 0 && hour <= 6:
				ticketsCountByTime["inicio_manha"]++
			case hour >= 7 && hour <= 12:
				ticketsCountByTime["manha"]++
			case hour >= 13 && hour <= 19:
				ticketsCountByTime["tarde"]++
			case hour >= 20 && hour <= 23:
				ticketsCountByTime["noite"]++
			default:
				fmt.Printf("Aviso: Hora do ticket fora do range esperado (0-23)")
			}
		}
	}

	return ticketsCountByTime, nil
}

func (t *repository) AverageDestination(destination string, time int) (int, error) {

	totalTickets := t.data
	var ticketsFound int
	var totalTicketsDay int
	for _, ticket := range totalTickets {
		hourString := ticket.Arrival
		hourParts := strings.Split(hourString, ":")
		hour, err := strconv.Atoi(hourParts[0])
		if err != nil {
			fmt.Errorf("Error: Houve um erro ao tentar converter a hora")
		}
		if ticket.Destination == destination {
			if time == hour {
				ticketsFound++
			}
			totalTicketsDay++
		}
	}

	result := float64(ticketsFound) / float64(totalTicketsDay)
	return int(result * 100), nil
}
