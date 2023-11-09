package main

// observer interface
type IObserver interface {
	update(string)
	getID() string
}
