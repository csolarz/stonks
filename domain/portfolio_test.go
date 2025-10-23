package domain

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper: comparar floats con margen de error
func almostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestPortfolioRebalanceMetaRises(t *testing.T) {
	// Crear acciones
	meta := NewStock("META", 1.0)
	googl := NewStock("GOOGL", 1.0)

	// Crear portafolio inicial de $100 con distribución 60/40
	p := NewPortfolio(100.0)
	p.AddStock(meta, 0.6)  // 60%
	p.AddStock(googl, 0.4) // 40%

	// Rebalance inicial
	p.Rebalance()

	// Validar condiciones iniciales
	total := p.TotalValue()
	if !almostEqual(total, 100.0, 0.001) {
		t.Errorf("Expected total 100.0, got %.2f", total)
	}

	// META sube a $2
	meta.SetPrice(2.0)

	// Validar nuevo valor total
	expectedTotal := 160.0
	newTotal := p.TotalValue()
	if !almostEqual(newTotal, expectedTotal, 0.001) {
		t.Errorf("Expected total %.2f, got %.2f", expectedTotal, newTotal)
	}

	// Validar redistribución 60/40
	metaValue := p.quantities["META"] * meta.Price()
	googlValue := p.quantities["GOOGL"] * googl.Price()

	metaPct := metaValue / newTotal
	googlPct := googlValue / newTotal

	if !almostEqual(metaPct, 0.6, 0.01) {
		t.Errorf("Expected META to be 60%%, got %.2f%%", metaPct*100)
	}

	if !almostEqual(googlPct, 0.4, 0.01) {
		t.Errorf("Expected GOOGL to be 40%%, got %.2f%%", googlPct*100)
	}

	// Validar que no haya dinero "perdido"
	sum := metaValue + googlValue + p.cash
	if !almostEqual(sum, newTotal, 0.001) {
		t.Errorf("Portfolio accounting error: expected %.2f, got %.2f", newTotal, sum)
	}
}

func TestPortfolioThreeStocksUpAndDown(t *testing.T) {
	// Crear acciones
	meta := NewStock("META", 100)
	googl := NewStock("GOOGL", 100)
	msft := NewStock("MSFT", 100)

	// Crear portafolio de $1,000 con 50/30/20
	p := NewPortfolio(1000)
	p.AddStock(meta, 0.5)
	p.AddStock(googl, 0.3)
	p.AddStock(msft, 0.2)

	p.Rebalance()

	assert.EqualValues(t, 5, p.quantities["META"])
	assert.EqualValues(t, 3, p.quantities["GOOGL"])
	assert.EqualValues(t, 2, p.quantities["MSFT"])
	assert.EqualValues(t, 0, p.cash)

	initialTotal := p.TotalValue()
	if !almostEqual(initialTotal, 1000, 0.001) {
		t.Errorf("Expected initial total 1000, got %.2f", initialTotal)
	}

	// ---- Caso 1: META sube de 100 a 150 ----
	meta.SetPrice(150)
	totalAfterMetaRise := p.TotalValue()

	if totalAfterMetaRise <= initialTotal {
		t.Errorf("Expected portfolio to increase after META rises")
	}

	// Verificar proporciones (aprox. 50%, 30%, 20%)
	metaValue := p.quantities["META"] * meta.Price()
	googlValue := p.quantities["GOOGL"] * googl.Price()
	msftValue := p.quantities["MSFT"] * msft.Price()
	total1 := p.TotalValue()

	metaPct := metaValue / total1
	googlPct := googlValue / total1
	msftPct := msftValue / total1

	assert.EqualValues(t, 0, p.cash)

	if !almostEqual(metaPct, 0.5, 0.02) {
		t.Errorf("Expected META ~50%%, got %.2f%%", metaPct*100)
	}
	if !almostEqual(googlPct, 0.3, 0.02) {
		t.Errorf("Expected GOOGL ~30%%, got %.2f%%", googlPct*100)
	}
	if !almostEqual(msftPct, 0.2, 0.02) {
		t.Errorf("Expected MSFT ~20%%, got %.2f%%", msftPct*100)
	}

	// ---- Caso 2: GOOGL baja de 100 a 80 ----
	googl.SetPrice(80)
	totalAfterGooglDrop := p.TotalValue()

	assert.EqualValues(t, 0, p.cash)

	if totalAfterGooglDrop >= totalAfterMetaRise {
		t.Errorf("Expected total to decrease after GOOGL drops")
	}

	// Recalcular proporciones
	metaValue = p.quantities["META"] * meta.Price()
	googlValue = p.quantities["GOOGL"] * googl.Price()
	msftValue = p.quantities["MSFT"] * msft.Price()
	total2 := p.TotalValue()

	metaPct = metaValue / total2
	googlPct = googlValue / total2
	msftPct = msftValue / total2

	if !almostEqual(metaPct, 0.5, 0.02) {
		t.Errorf("After drop, expected META ~50%%, got %.2f%%", metaPct*100)
	}
	if !almostEqual(googlPct, 0.3, 0.02) {
		t.Errorf("After drop, expected GOOGL ~30%%, got %.2f%%", googlPct*100)
	}
	if !almostEqual(msftPct, 0.2, 0.02) {
		t.Errorf("After drop, expected MSFT ~20%%, got %.2f%%", msftPct*100)
	}

	assert.EqualValues(t, 0, p.cash)

	// Validar que no haya dinero perdido
	sum := metaValue + googlValue + msftValue + p.cash
	if !almostEqual(sum, total2, 0.001) {
		t.Errorf("Accounting error: expected total %.3f, got sum %.3f", total2, sum)
	}
}

func TestPortfolio_AddStock(t *testing.T) {
	tests := []struct {
		name          string  // description of this test case
		initialAmount float64 // initial cash in portfolio
		// Named input parameters for target function.
		s          *Stock
		percentage float64
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name:          "valid percentage",
			initialAmount: 1000,
			s:             NewStock("AAPL", 150),
			percentage:    0.5,
			wantErr:       false,
		},
		{
			name:          "percentage zero",
			initialAmount: 1000,
			s:             NewStock("AAPL", 150),
			percentage:    0.0,
			wantErr:       true,
		},
		{
			name:          "percentage negative",
			initialAmount: 1000,
			s:             NewStock("AAPL", 150),
			percentage:    -0.2,
			wantErr:       true,
		},
		{
			name:          "percentage over rated",
			initialAmount: 1000,
			s:             NewStock("AAPL", 150),
			percentage:    1.5,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPortfolio(tt.initialAmount)
			gotErr := p.AddStock(tt.s, tt.percentage)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("AddStock() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("AddStock() succeeded unexpectedly")
			}
		})
	}
}
