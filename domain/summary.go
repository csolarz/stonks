package domain

import (
	"fmt"
)

func ShowSummary(p Portfolio) {
	total := p.TotalValue()
	fmt.Println("\n📊 Current portfolio summary:")
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("%-10s %-10s %-10s %-12s %-12s\n", "Stock", "Price", "Qty", "Value", "Share")

	for name, stock := range p.stocks {
		value := stock.Price() * p.quantities[name]
		share := (value / total) * 100
		fmt.Printf("%-10s %-10.3f %-10.3f %-12.3f %-10.3f%%\n",
			name, stock.Price(), p.quantities[name], value, share)
	}

	fmt.Printf("\n💵 Cash available: %.3ff\n", p.cash)
	fmt.Printf("💼 Total portfolio value: %.3ff\n", total)
	fmt.Println("===============================================================")
}
