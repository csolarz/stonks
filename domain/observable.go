package domain

// Updater defines the observer interface for stock updates
type Updater interface {
	Update(a *Stock)
}

// Subject defines the subject interface for managing observers
type Subject interface {
	Subscribe(o Updater)
	Unsubscribe(o Updater)
	Notify()
}
