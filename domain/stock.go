package domain

type Stock struct {
	nombre       string
	valor        float64
	observadores []Updater
}

func NewStock(nombre string, valor float64) *Stock {
	return &Stock{
		nombre: nombre,
		valor:  valor,
	}
}

func (a *Stock) Name() string {
	return a.nombre
}

func (a *Stock) Price() float64 {
	return a.valor
}

// SetValor actualiza el precio y notifica a los observadores
func (a *Stock) SetValor(valor float64) {
	a.valor = valor
	a.Notificar()
}

// Métodos de sujeto
func (a *Stock) Register(o Updater) {
	a.observadores = append(a.observadores, o)
}

func (a *Stock) Remover(o Updater) {
	for i, obs := range a.observadores {
		if obs == o {
			a.observadores = append(a.observadores[:i], a.observadores[i+1:]...)
			break
		}
	}
}

func (a *Stock) Notificar() {
	for _, obs := range a.observadores {
		obs.Update(a)
	}
}
