package main

import (
	"fmt"

	"github.com/csolarz/stonks/domain"
)

func main() {
	metaStocks := domain.NewStock("META", 1)
	googleStocks := domain.NewStock("GOOGL", 1)

	port := domain.NewPortfolio(100)
	port.AddStock(metaStocks, 0.6)   // 60% META
	port.AddStock(googleStocks, 0.4) // 40% GOOGL

	fmt.Println("Portafolio inicial:")
	port.Rebalance()

	// Cambios de precio
	fmt.Println("Subida de META a 2")
	metaStocks.SetValor(2) // Subida -> debe vender META y comprar GOOGL

	fmt.Println("Subida de META a 1000")
	metaStocks.SetValor(1000) // Subida -> debe vender META y comprar GOOGL
}
