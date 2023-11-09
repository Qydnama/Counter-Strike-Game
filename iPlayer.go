package main

// observable interface
// factory interface
// strategy interface
type IPlayer interface {
	SetTeam(name string)
	GetTeam() string
	SetWeaponStrategy(strategy IWeapon)
	UseWeapon()
	GetWeaponStrategy() IWeapon
	SetEquipmentStrategy(strategy IEquipment)
	UseEquipment()
	register(observer IObserver)
	deregister(observer IObserver)
	notifyAll()
	updateAvailability()
}
