package ticket_csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/luizfeleal/desafio-go-bases/internal/tickets/domain"
)

func GetTicketsCsvContent(path string) map[int][]domain.Ticket {
	file, err := os.Open(path)

	if err != nil {
		fmt.Errorf("An error ocurrs trying to get open the file")
		return nil
	}

	defer file.Close()

	reader := csv.NewReader(file)

	basicData := make(map[int][]domain.Ticket)
	for {
		var ticket domain.Ticket

		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Errorf("An error ocurs trying to read the csv file")
			return nil
		}

		if err = ticket.NormalizedTicket(record); err != nil {
			fmt.Errorf("An error ocurs trying to make a normalized ticket")
			return nil
		}

		basicData[ticket.Id] = append(basicData[ticket.Id], ticket)
	}

	return basicData
}
