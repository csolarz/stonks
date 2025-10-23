package main

import (
	"fmt"

	"github.com/csolarz/stonks/domain"
)

func main() {
	stockA := domain.NewStock("AAPL", 1)
	stockB := domain.NewStock("GOOGL", 1)

	// Crear portafolio con 2 acciones
	port := domain.NewPortfolio(100)
	port.AddStock(stockA, 0.6) // 60%
	port.AddStock(stockB, 0.4) // 40%

	// Initial rebalance for baseline
	port.Rebalance()

	fmt.Println("🔹 Initial portfolio state:")
	domain.ShowSummary(*port)

	// Cambios simulados
	fmt.Println("\n=== AAPL rises to 2 ===")
	stockA.SetPrice(2)
	domain.ShowSummary(*port)

	fmt.Println("\n=== AAPL rises to 200 ===")
	stockA.SetPrice(200)
	domain.ShowSummary(*port)
}
