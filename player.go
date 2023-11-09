package main

import (
	"fmt"
)

type Player struct {
	team         string
	weapon       IWeapon
	equipment    IEquipment
	observerList []IObserver
	inStock      bool
}

func (p *Player) SetTeam(name string) {
	p.team = name
}

func (p *Player) GetTeam() string {
	return p.team
}

func (p *Player) SetWeaponStrategy(strategy IWeapon) {
	p.weapon = strategy
}

func (p *Player) UseWeapon() {
	if p.weapon != nil {
		p.weapon.UseWeapon()
	} else {
		fmt.Println("No weapon assigned to the player.")
	}
}
func (p *Player) GetWeaponStrategy() IWeapon {
	return p.weapon
}

func (p *Player) SetEquipmentStrategy(strategy IEquipment) {
	p.equipment = strategy
}

func (p *Player) UseEquipment() {
	if p.equipment != nil {
		p.equipment.UseEquipment()
	} else {
		fmt.Println("No equipment assigned to the player.")
	}
}

func (i *Player) updateAvailability() {
	fmt.Printf("Player %s is sending information\n", i.team)
	i.inStock = true
	i.notifyAll()
}
func (i *Player) register(o IObserver) {
	i.observerList = append(i.observerList, o)
}

func (i *Player) deregister(o IObserver) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *Player) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.team)
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
