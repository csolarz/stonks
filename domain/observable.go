package domain

// Updater representa al observador que reacciona a los cambios del sujeto
type Updater interface {
	Update(s *Stock)
}

// Subject representa al sujeto que notifica a sus observadores
type Subject interface {
	Subscribe(u Updater)
	Unsubscribe(u Updater)
	Notify()
}
