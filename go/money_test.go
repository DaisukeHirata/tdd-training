package main

import (
	s "tdd/stocks"
	"testing"
)

func assertEqual(t *testing.T, expected, actual interface{}) {
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
	portFolioInDollars, _ = portFolio.Evaluate("USD")

	assertEqual(t, fifteenDollars, portFolioInDollars)
}

func TestAddtionOfDollarsAndEuros(t *testing.T) {
	var portFolio s.PortFolio

	fiveDollars := s.NewMoney(5, "USD")
	tenEuros := s.NewMoney(10, "EUR")

	portFolio = portFolio.Add(fiveDollars)
	portFolio = portFolio.Add(tenEuros)

	expectedValue := s.NewMoney(17, "USD")
	actualValue, _ := portFolio.Evaluate("USD")

	assertEqual(t, expectedValue, actualValue)
}

func TestAdditionOfDollarsAndWons(t *testing.T) {
	var portFolio s.PortFolio

	oneDollar := s.NewMoney(1, "USD")
	elevenHundredWon := s.NewMoney(1100, "KRW")

	portFolio = portFolio.Add(oneDollar)
	portFolio = portFolio.Add(elevenHundredWon)

	expectedValue := s.NewMoney(2200, "KRW")
	actualValue, _ := portFolio.Evaluate("KRW")

	assertEqual(t, expectedValue, actualValue)
}

func TestAdditionWithMultipleMissingExchangeRates(t *testing.T) {
	var portFolio s.PortFolio

	oneDollar := s.NewMoney(1, "USD")
	oneEuro := s.NewMoney(1, "EUR")
	oneWon := s.NewMoney(1, "KRW")

	portFolio = portFolio.Add(oneDollar)
	portFolio = portFolio.Add(oneEuro)
	portFolio = portFolio.Add(oneWon)

	expectedErrorMessage := "Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid,]"
	_, actualError := portFolio.Evaluate("Kalganid")

	assertEqual(t, expectedErrorMessage, actualError.Error())
}
