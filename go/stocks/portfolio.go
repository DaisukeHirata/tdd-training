package stocks

type PortFolio []Money

func (p PortFolio) Add(money Money) PortFolio {
	p = append(p, money)
	return p
}

func (p PortFolio) Evaluate(currency string) Money {
	total := 0.0
	for _, money := range p {
		total += convert(money, currency)
	}
	return Money{amount: total, currency: currency}
}

func convert(money Money, currency string) float64 {
	eurToUsd := 1.2
	if money.currency == currency {
		return money.amount
	}
	return money.amount * eurToUsd
}
