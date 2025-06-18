package domain

import (
	"strconv"
)

type Ticket struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Destination string `json:"destination"`
	Arrival     string `json:"arrival"`
	Price       int    `json:"price"`
}

func (t *Ticket) NormalizedTicket(s []string) error {
	id, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	t.Id = id
	t.Name = s[1]
	t.Email = s[2]
	t.Destination = s[3]
	t.Arrival = s[4]

	price, err := strconv.Atoi(s[5])
	if err != nil {
		return err
	}
	t.Price = price
	return nil
}
