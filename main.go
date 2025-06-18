package main

import (
	"fmt"

	"github.com/luizfeleal/desafio-go-bases/internal/tickets"
)

type desafioFinalServer struct {
	repository tickets.Ticket
}

var server = desafioFinalServer{
	repository: tickets.NewRepository(),
}

func main() {
	country := "china"
	server.repository.GetTotalTickets(country)
	total, err := server.repository.GetCountryByPeriod(country)
	if err != nil {
		fmt.Errorf("deu erro")
	}
	fmt.Println(total)
	server.repository.AverageDestination(country, 1)
}
