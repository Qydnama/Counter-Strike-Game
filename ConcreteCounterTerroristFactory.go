package main

type CounterTerroristFactory struct {
}

func (t *CounterTerroristFactory) createRifle() IWeapon {
	return &M4{}
}

func (t *CounterTerroristFactory) createPistol() IWeapon {
	return &USP{}
}

func (t *CounterTerroristFactory) createKnife() IWeapon {
	return &KnifeCounterTerrorist{}
}

func (t *CounterTerroristFactory) createEquipment() IEquipment {
	return &DefuseKitStrategy{Bomb: globalBomb}
}
