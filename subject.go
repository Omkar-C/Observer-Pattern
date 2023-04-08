package main

type Subject interface {
	Attach(Observer)
	Notify()
}

type SubjectImpl struct {
	observers []Observer
}

func (s *SubjectImpl) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *SubjectImpl) Notify(data Data) {
	for _, o := range s.observers {
		o.Update(data)
	}
}
