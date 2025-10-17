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
func (p *Portfolio) AddStock(s *Stock, percentage float64) {
	p.stocks[s.Name()] = s
	p.allocations[s.Name()] = &AllocatedStock{name: s.Name(), rate: percentage}

	// Compra inicial proporcional
	investment := p.cash * percentage
	qty := investment / s.Price()
	p.quantities[s.Name()] = qty
	p.cash -= investment
	s.Subscribe(p)
}

// Método Updater
func (p *Portfolio) Update(s *Stock) {
	fmt.Printf("\n📈 Change detected in %s -> new price: %.2f\n", s.Name(), s.Price())
	p.Rebalance()
}

// ------------------ Lógica de rebalanceo ------------------

func (p *Portfolio) totalValue() float64 {
	total := p.cash
	for name, stock := range p.stocks {
		total += p.quantities[name] * stock.Price()
	}
	return total
}

func (p *Portfolio) Rebalance() {
	total := p.totalValue()
	fmt.Printf("💰 Total portfolio value: %.2f (Cash: %.2f)\n", total, p.cash)

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
			if p.cash >= diff {
				p.quantities[name] += qty
				p.cash -= diff
				fmt.Printf("🟢 Buying %.2f of %s (%.2f USD)\n", qty, name, diff)
			} else {
				fmt.Printf("⚠️ Not enough cash to buy %.2f of %s\n", qty, name)
			}
		} else if diff < 0 {
			// Vender
			qty := (-diff) / stock.Price()
			if p.quantities[name] >= qty {
				p.quantities[name] -= qty
				p.cash += -diff
				fmt.Printf("🔴 Selling %.2f of %s (%.2f USD)\n", qty, name, -diff)
			}
		}
	}

	p.ShowSummary()
}

func (p *Portfolio) ShowSummary() {
	total := p.totalValue()
	fmt.Println("\n📊 Current portfolio summary:")
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("%-10s %-10s %-10s %-12s %-12s\n", "Stock", "Price", "Qty", "Value", "Share")

	for name, stock := range p.stocks {
		value := stock.Price() * p.quantities[name]
		share := (value / total) * 100
		fmt.Printf("%-10s %-10.2f %-10.2f %-12.2f %-10.2f%%\n",
			name, stock.Price(), p.quantities[name], value, share)
	}

	fmt.Printf("\n💵 Cash available: %.2f\n", p.cash)
	fmt.Printf("💼 Total portfolio value: %.2f\n", total)
	fmt.Println("===============================================================")
}
