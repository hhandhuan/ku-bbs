package remind

import "github.com/huhaophp/hblog/internal/subject"

func New() *remindSubject {
	return &remindSubject{
		observers: make([]subject.Observer, 0),
	}
}

type remindSubject struct {
	observers []subject.Observer
}

func (o *remindSubject) Attach(observer subject.Observer) {
	o.observers = append(o.observers, observer)
}

func (o *remindSubject) Notify() {
	for _, s := range o.observers {
		s.Update()
	}
}
