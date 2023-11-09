package main

type IGameInstance interface {
	register(observer IObserver)
	deregister(observer IObserver)
	notifyAll(comment string)
	updateAvailability(s string)
}
