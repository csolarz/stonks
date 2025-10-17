package domain

// Updater representa al observador que reacciona a los cambios del sujeto
type Updater interface {
	Update(a *Stock)
}

// Subject representa al sujeto que notifica a sus observadores
type Subject interface {
	Subscribe(o Updater)
	Unsubscribe(o Updater)
	Notify()
}
