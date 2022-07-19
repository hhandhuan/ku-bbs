package subject

type Subject interface {
	Attach(Observer)
	Notify()
}

type Observer interface {
	Update()
}
