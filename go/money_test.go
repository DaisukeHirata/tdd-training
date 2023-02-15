package main

import (
	s "tdd/stocks"
	"testing"
)

func assertEqual(t *testing.T, expected, actual s.Money) {
	if expected != actual {
		t.Errorf("Expected %+v, got: %+v", expected, actual)
	}
}

func TestMultiplication(t *testing.T) {
	tenEuros := s.NewMoney(10, "EUR")
	actualResult := tenEuros.Times(2)
	expectedResult := s.NewMoney(20, "EUR")
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := s.NewMoney(4002, "KRW")
	actualResult := originalMoney.Divide(4)
	expectedResult := s.NewMoney(1000.5, "KRW")
	assertEqual(t, expectedResult, actualResult)
}

func TestAddition(t *testing.T) {
	var portFolio s.PortFolio
	var portFolioInDollars s.Money

	fiveDollars := s.NewMoney(5, "USD")
	tenDollars := s.NewMoney(10, "USD")
	fifteenDollars := s.NewMoney(15, "USD")

	portFolio = portFolio.Add(fiveDollars)
	portFolio = portFolio.Add(tenDollars)
	portFolioInDollars = portFolio.Evaluate("USD")

	assertEqual(t, fifteenDollars, portFolioInDollars)
}

func TestAddtionOfDollarsAndEuros(t *testing.T) {
	var portFolio s.PortFolio

	fiveDollars := s.NewMoney(5, "USD")
	tenEuros := s.NewMoney(10, "EUR")

	portFolio = portFolio.Add(fiveDollars)
	portFolio = portFolio.Add(tenEuros)

	expectedValue := s.NewMoney(17, "USD")
	actualValue := portFolio.Evaluate("USD")

	assertEqual(t, expectedValue, actualValue)
}
