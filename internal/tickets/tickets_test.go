package tickets_test

import (
	"fmt"
	"testing"

	"github.com/luizfeleal/desafio-go-bases/internal/tickets"
)

func TestGetTotalTickets(t *testing.T) {
	repository := tickets.NewRepository(
		tickets.WithFilePath("./tickets.csv"),
	)
	expectedResult := 29
	expectedCase := "China"

	result, err := repository.GetTotalTickets(expectedCase)

	if err != nil {
		t.Error("Error trying to get the total Tickets")
	}

	if expectedResult != result {
		t.Errorf("GetTotalTickets() function gave the result -%v, but the expected result is %v", result, expectedResult)
	}

}

func TestGetCountryByPeriod(t *testing.T) {
	repository := tickets.NewRepository(
		tickets.WithFilePath("./tickets.csv"),
	)
	expectedResultManha := 10
	expectedCase := "China"

	result, err := repository.GetCountryByPeriod(expectedCase)

	fmt.Println(result)
	if err != nil {
		t.Error("Error trying to get the total Tickets")
	}

	if expectedResultManha != result["manha"] {
		t.Errorf("GetCountryByPeriod() function gave the result -%v, but the expected result is %v", result, expectedResultManha)
	}
}

func TestAverageDestination(t *testing.T) {
	repository := tickets.NewRepository(
		tickets.WithFilePath("./tickets.csv"),
	)
	expectedResult := 6
	expectedCaseDestination := "China"
	expectedCaseTime := 17

	result, err := repository.AverageDestination(expectedCaseDestination, expectedCaseTime)
	fmt.Println("cehguei")
	fmt.Println(result)
	if err != nil {
		t.Error("Error trying to get the total Tickets")
	}

	if expectedResult != result {
		t.Errorf("GetCountryByPeriod() function gave the result -%v, but the expected result is %v", result, expectedResult)
	}
}
