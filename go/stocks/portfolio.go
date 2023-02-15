package stocks

type PortFolio []Money

func (p PortFolio) Add(money Money) PortFolio {
	p = append(p, money)
	return p
}

func (p PortFolio) Evaluate(currency string) Money {
	total := 0.0
	for _, money := range p {
		total += money.amount
	}
	return Money{amount: total, currency: currency}
}
