package main

import "fmt"

type Observer interface {
	Update(Data)
}

type ObserverImpl struct {
	ID int
}

func (o *ObserverImpl) Update(data Data) {
	fmt.Println("Observer", o.ID, "received data", data.name)
}
