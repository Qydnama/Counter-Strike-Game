package main

import (
	"fmt"
)

var gameInstance = GetGameInstance()

func main() {
	var input int
	observerFirst := &User{id: "1"}
	observerSecond := &User{id: "2"}
	var terrorist IPlayer
	var cterrorist IPlayer
	fmt.Println("You are playing bomb Defuse game!")
	fmt.Print("First you are playing for Terrorist")

	terrorist, err := GetTerrorist()
	if err != nil {
		panic("Terrorist can't be created")
	}
	bomb := &BombStrategy{}
	terrorist.SetEquipmentStrategy(bomb)
	fmt.Println(terrorist.GetTeam())
	terrorist.register(observerFirst)
	terrorist.register(observerSecond)
	terrorist.updateAvailability()

	//to check if players exists
	GetGameInstance()

	fmt.Println("Choose a weapon (1)-Knife (2)-Gun (for Terrorist)")
	fmt.Print("Your Input: ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		knifeStrategy := &KnifeStrategy{}
		terrorist.SetWeaponStrategy(knifeStrategy)
	case 2:
		gunStrategy := &GunStrategy{}
		terrorist.SetWeaponStrategy(gunStrategy)
	default:
		panic("Wrong weapon type choosed")
	}

	cterrorist, err = GetCounterTerrorist()
	if err != nil {
		panic("Counter-Terrorist can't be created")
	}

	fmt.Println(cterrorist.GetTeam())
	cterrorist.register(observerFirst)
	cterrorist.register(observerSecond)
	cterrorist.updateAvailability()
	cterrorist.GetTeam()

	fmt.Println("Choose a weapon (1)-Knife (2)-Gun (for Counter-Terrorist)")
	fmt.Print("Your Input: ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		knifeStrategy := &KnifeStrategy{}
		cterrorist.SetWeaponStrategy(knifeStrategy)
	case 2:
		gunStrategy := &GunStrategy{}
		cterrorist.SetWeaponStrategy(gunStrategy)
	default:
		panic("Wrong weapon type choosed")
	}

	fmt.Println("Buy a Defuse Kit?:  (1)-Yes (2)-No  (for Counter-Terrorist)")
	fmt.Print("Your Input: ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		defuseKit := &DefuseKitStrategy{Exist: true}
		cterrorist.SetEquipmentStrategy(defuseKit)
	case 2:
		defuseKit := &DefuseKitStrategy{Exist: false}
		cterrorist.SetEquipmentStrategy(defuseKit)
	default:
		panic("Wrong symbol entered")
	}

	printDetailss(terrorist)

}

func printDetailss(p IPlayer) {
	fmt.Println("Someone shooting at you!")
	p.UseWeapon()
	fmt.Println("You are dead!")
}
