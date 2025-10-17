package domain

import "fmt"

type Portfolio struct {
	amount          float64
	allocatedStocks map[string]*AllocatedStock
	stocks          map[string]*Stock
	quantity        map[string]float64 // Cantidad de cada acción
}

func NewPortfolio(amount float64) *Portfolio {
	return &Portfolio{
		amount:          amount,
		allocatedStocks: make(map[string]*AllocatedStock),
		stocks:          make(map[string]*Stock),
		quantity:        make(map[string]float64),
	}
}

func (p *Portfolio) AddStock(stock *Stock, rate float64) {
	p.stocks[stock.Name()] = stock
	p.allocatedStocks[stock.Name()] = &AllocatedStock{
		name: stock.Name(),
		rate: rate,
	}
	// Inicialmente comprar acciones según porcentaje
	p.quantity[stock.Name()] = (p.amount * rate) / stock.Price()
	stock.Register(p)
}

// Método Observador: rebalanceo
func (p *Portfolio) Update(a *Stock) {
	fmt.Printf("Precio de %s cambió a %.2f\n", a.Name(), a.Price())
	p.Rebalance()
}

// Rebalancea vendiendo/comprando para mantener porcentajes
func (p *Portfolio) Rebalance() {
	total := 0.0
	// Calcular valor total del portafolio actual
	for name, stock := range p.stocks {
		total += p.quantity[name] * stock.Price()
	}

	fmt.Printf("Valor total del portafolio: %.2f\n", total)

	for name, allocated := range p.allocatedStocks {
		stock := p.stocks[name]

		desiredAmount := total * allocated.rate
		desiredQuantity := desiredAmount / stock.Price()
		change := desiredQuantity - p.quantity[name]

		if change > 0 {
			fmt.Printf("Buy %.2f of %s\n", change, name)
		} else if change < 0 {
			fmt.Printf("Sell %.2f of %s\n", -change, name)
		}
		// Actualizar cantidad real
		p.quantity[name] = desiredQuantity
	}

	fmt.Println("---")
}
