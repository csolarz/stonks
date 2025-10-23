package domain

import "fmt"

type Portfolio struct {
	cash        float64
	allocations map[string]*AllocatedStock
	stocks      map[string]*Stock
	quantities  map[string]float64
}

func NewPortfolio(initialAmount float64) *Portfolio {
	return &Portfolio{
		cash:        initialAmount,
		allocations: make(map[string]*AllocatedStock),
		stocks:      make(map[string]*Stock),
		quantities:  make(map[string]float64),
	}
}

// Agrega una acción con su porcentaje al portafolio
func (p *Portfolio) AddStock(s *Stock, percentage float64) error {
	if percentage <= 0 || percentage > 1 || p.cash == 0 {
		return fmt.Errorf("percentage or cash invalid")
	}

	p.stocks[s.Name()] = s
	p.allocations[s.Name()] = &AllocatedStock{
		name: s.Name(),
		rate: percentage,
	}

	s.Subscribe(p)

	return nil
}

// Método Updater
func (p *Portfolio) Update(s any) {
	stock, ok := s.(*Stock)
	if !ok {
		return
	}

	fmt.Printf("\n📈 Change detected in %s -> new price: %.3ff\n", stock.Name(), stock.Price())
	p.Rebalance()
}

// ------------------ Lógica de rebalanceo ------------------

func (p *Portfolio) TotalValue() float64 {
	total := p.cash
	for name, stock := range p.stocks {
		total += p.quantities[name] * stock.Price()
	}
	return total
}

func (p *Portfolio) Rebalance() {
	total := p.TotalValue()
	fmt.Printf("💰 Total portfolio value: %.3ff (Cash: %.3ff)\n", total, p.cash)

	// 1️⃣ Calcular cuánto debería tener cada acción
	targetValues := make(map[string]float64)
	for name, alloc := range p.allocations {
		targetValues[name] = total * alloc.rate
	}

	// 2️⃣ Calcular diferencias y ajustar
	for name, stock := range p.stocks {
		currentValue := p.quantities[name] * stock.Price()
		targetValue := targetValues[name]
		diff := targetValue - currentValue

		if diff > 0 {
			// Comprar
			qty := diff / stock.Price()
			p.quantities[name] += qty
			p.cash -= diff
			fmt.Printf("🟢 Buying %.3ff of %s (%.3ff USD)\n", qty, name, diff)
		} else if diff < 0 {
			// Vender
			qty := (-diff) / stock.Price()
			p.quantities[name] -= qty
			p.cash += -diff
			fmt.Printf("🔴 Selling %.3ff of %s (%.3ff USD)\n", qty, name, -diff)
		}
	}
}
