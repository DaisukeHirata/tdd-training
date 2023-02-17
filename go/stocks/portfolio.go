package stocks

import "errors"

type PortFolio []Money

func (p PortFolio) Add(money Money) PortFolio {
	p = append(p, money)
	return p
}

func (p PortFolio) Evaluate(bank Bank, currency string) (*Money, error) {
	total := 0.0
	failedConversions := make([]string, 0)
	for _, money := range p {
		if convertedACurrency, err := bank.Convert(money, currency); err == nil {
			total += convertedACurrency.amount
		} else {
			failedConversions = append(failedConversions, money.currency+"->"+currency)
		}
	}
	if len(failedConversions) == 0 {
		totalMoney := NewMoney(total, currency)
		return &totalMoney, nil
	}
	failures := "["
	for _, failure := range failedConversions {
		failures += failure + ","
	}
	failures += "]"
	return nil, errors.New("Missing exchange rate(s):" + failures)
}
