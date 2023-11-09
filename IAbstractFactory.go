package main

import "fmt"

type IAbstractFactory interface {
	createRifle() IWeapon
	createPistol() IWeapon
	createKnife() IWeapon
	createEquipment() IEquipment
}

func GetAbstractFactory(brand string) (IAbstractFactory, error) {
	if brand == "terrorist" {
		return &TerroristFactory{}, nil
	}

	if brand == "counterTerrorist" {
		return &CounterTerroristFactory{}, nil
	}

	return nil, fmt.Errorf("Wrong Input")
}
