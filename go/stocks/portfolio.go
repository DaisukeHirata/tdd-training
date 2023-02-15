package stocks

import "errors"

type PortFolio []Money

func (p PortFolio) Add(money Money) PortFolio {
	p = append(p, money)
	return p
}

func (p PortFolio) Evaluate(currency string) (Money, error) {
	total := 0.0
	failedConversions := make([]string, 0)
	for _, money := range p {
		if convertedAmount, ok := convert(money, currency); ok {
			total += convertedAmount
		} else {
			failedConversions = append(failedConversions, money.currency+"->"+currency)
		}
	}
	if len(failedConversions) == 0 {
		return NewMoney(total, currency), nil
	}
	failures := "["
	for _, failure := range failedConversions {
		failures += failure + ","
	}
	failures += "]"
	return NewMoney(0, ""), errors.New("Missing exchange rate(s):" + failures)
}

func convert(money Money, currency string) (float64, bool) {
	exchangeRates := map[string]float64{
		"EUR->USD": 1.2,
		"USD->KRW": 1100,
	}
	if money.currency == currency {
		return money.amount, true
	}
	key := money.currency + "->" + currency
	rate, ok := exchangeRates[key]
	return money.amount * rate, ok
}
