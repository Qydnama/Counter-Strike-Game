package main

// factory interface
// strategy interface
type IPlayer interface {
	GetName() string
	GetTeam() string
	SetWeaponStrategy(strategy IWeapon)
	Kill() (string, bool)
	GetWeaponStrategy() IWeapon
	SetEquipmentStrategy(strategy IEquipment)
	UseEquipment()
}
