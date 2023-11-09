package main

type GameInstance struct {
	Players      []IPlayer
	observerList []IObserver
}

func (i *GameInstance) register(o IObserver) {
	i.observerList = append(i.observerList, o)
}

func (i *GameInstance) deregister(o IObserver) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *GameInstance) notifyAll(comment string) {
	for _, observer := range i.observerList {
		observer.update(comment)
	}
}

func removeFromslice(observerList []IObserver, observerToRemove IObserver) []IObserver {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
