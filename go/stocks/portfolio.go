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
	exchangeRates := map[string]float64{
		"EUR->USD": 1.2,
		"USD->KRW": 1100,
	}
	if money.currency == currency {
		return money.amount
	}
	key := money.currency + "->" + currency
	return money.amount * exchangeRates[key]
}
