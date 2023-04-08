package main

func main() {
	subject := SubjectImpl{}
	
	observer1 := ObserverImpl{ID: 1}
	observer2 := ObserverImpl{ID: 2}
	observer3 := ObserverImpl{ID: 3}

	subject.Attach(&observer1)
	subject.Attach(&observer2)
	subject.Attach(&observer3)

	subject.Notify(Data{name: "Hello World"})
	subject.Notify(Data{name: "Hello World 2"})
}