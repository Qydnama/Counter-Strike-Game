package main

var globalBomb *Bomb

type TerroristFactory struct {
}

func (t *TerroristFactory) createRifle() IWeapon {
	return &AK47{}
}

func (t *TerroristFactory) createPistol() IWeapon {
	return &Glock{}
}

func (t *TerroristFactory) createKnife() IWeapon {
	return &Knife{}
}

func (t *TerroristFactory) createEquipment() IEquipment {
	globalBomb = &Bomb{}
	return globalBomb
}
