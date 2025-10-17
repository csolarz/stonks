package main

import (
	"fmt"

	"github.com/csolarz/stonks/domain"
)

func main() {
	stockA := domain.NewStock("AAPL", 1)
	stockB := domain.NewStock("GOOGL", 1)
	//stockC := domain.NewStock("MSFT", 200)

	// Crear portafolio con 3 acciones
	port := domain.NewPortfolio(100)
	port.AddStock(stockA, 0.6) // 50%
	port.AddStock(stockB, 0.4) // 30%

	// Initial rebalance for baseline
	port.Rebalance()

	//port.AddStock(stockC, 0.2) // 20%

	fmt.Println("🔹 Initial portfolio state:")
	port.ShowSummary()

	// Cambios simulados
	fmt.Println("\n=== AAPL rises to 2 ===")
	stockA.SetPrice(2)
	port.ShowSummary()

	fmt.Println("\n=== AAPL rises to 200 ===")
	stockA.SetPrice(200)
	port.ShowSummary()

	//fmt.Println("\n=== MSFT rises to 250 ===")
	//stockC.SetPrice(250)
}
