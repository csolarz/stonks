package domain

import (
	"sync"

	"github.com/csolarz/stonks/observable"
)

type Stock struct {
	mu       sync.Mutex
	name     string
	price    float64
	updaters []observable.Updater
}

func NewStock(name string, price float64) *Stock {
	return &Stock{name: name, price: price, mu: sync.Mutex{}}
}

func (s *Stock) Name() string {
	return s.name
}

func (s *Stock) Price() float64 {
	return s.price
}

// SetPrice actualiza el precio y notifica a los observadores
func (s *Stock) SetPrice(price float64) {
	s.mu.Lock()
	s.price = price
	s.Notify()
	s.mu.Unlock()
}

// Métodos del Subject
func (s *Stock) Subscribe(u observable.Updater) {
	s.updaters = append(s.updaters, u)
}

func (s *Stock) Unsubscribe(u observable.Updater) {
	for i, obs := range s.updaters {
		if obs == u {
			s.updaters = append(s.updaters[:i], s.updaters[i+1:]...)
			break
		}
	}
}

func (s *Stock) Notify() {
	for _, obs := range s.updaters {
		obs.Update(s)
	}
}
