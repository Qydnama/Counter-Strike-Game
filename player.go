package main

type Player struct {
	team      string
	weapon    IWeapon
	equipment IEquipment
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

func (p *Player) Kill() (string, bool) {
	return p.weapon.Kill()
}

func (p *Player) GetWeaponStrategy() IWeapon {
	return p.weapon
}

func (p *Player) SetEquipmentStrategy(strategy IEquipment) {
	p.equipment = strategy
}

func (p *Player) UseEquipment() {
	p.equipment.UseEquipment()
}
