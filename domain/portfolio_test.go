package domain

import "testing"

func TestPortfolio_Rebalance(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		amount float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPortfolio(tt.amount)
			p.Rebalance()
		})
	}
}
