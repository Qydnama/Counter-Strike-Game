package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

var gameInstance = GetGameInstance()

func main() {
	var input int
	// Create Terrorist fabric by Using Abstract Factory

	afTerrorist, err := GetAbstractFactory("terrorist")
	if err != nil {
		panic("Terrorist factory can't create.")
	}
	// Create Counter-Terrorist fabric by Using Abstract Factory
	afCounterTerrorist, err := GetAbstractFactory("counterTerrorist")
	if err != nil {
		panic("Counter Terrorist factory can't create.")

	}

	// Users were born)
	observerFirst := &User{id: "1"}
	observerSecond := &User{id: "2"}

	// Users Added to the Match Like Observers
	gameInstance.register(observerFirst)
	gameInstance.register(observerSecond)

	// Creating Terrorist by Factory and Notify about it to Observers
	color.Set(color.FgYellow, color.BgHiCyan, color.Bold, color.Underline)
	fmt.Print("You are playing \"Bomb Defuse game!\"(BABAH:) ahahhaha)")
	color.Unset()
	fmt.Println("\n")
	var name string
	color.Set(color.FgMagenta)
	fmt.Print("Input Terrorist Name: ")
	fmt.Scan(&name)
	player1, err := GetPlayer("terrorist", name)
	if err != nil {
		panic("Terrorist can't be created")
	}

	// Creating Counter-Terrorist by Factory and Notify about it to Observers
	gameInstance.notifyAll(fmt.Sprintf("%s-%s was created.\n", player1.GetTeam(), player1.GetName()))
	fmt.Print("Input Counter-Terrorist Name: ")
	fmt.Scan(&name)
	player2, err := GetPlayer("counter-terrorist", name)
	if err != nil {
		panic("Counter-Terrorist can't be created")
	}
	gameInstance.notifyAll(fmt.Sprintf("%s-%s was created.\n", player2.GetTeam(), player2.GetName()))
	color.Unset()

	// Choosing Weapon for Terrorist
	//WEAPONTerror:
	//	for {
	color.Set(color.FgRed)
	//fmt.Print("Your Input: ")
	prompt := &survey.Select{
		Message: "Terrorist is choosing the weapon",
		Options: []string{"Rifle", "Pistol", "Knife"},
	}
	survey.AskOne(prompt, &input)
	color.Unset()

	switch input {
	case 0:
		player1.SetWeaponStrategy(afTerrorist.createRifle())
	case 1:
		player1.SetWeaponStrategy(afTerrorist.createPistol())
	case 2:
		player1.SetWeaponStrategy(afTerrorist.createKnife())
		//default:
		//	continue
	}
	//	break WEAPONTerror
	//}

	// Choosing Weapon for Counter-Terrorist
	//WEAPONCTerrorst:
	//	for {
	color.Set(color.FgBlue)
	prompt = &survey.Select{
		Message: "Counter Terrorist choose a weapon:",
		Options: []string{"Rifle", "Pistol", "Knife"},
	}
	survey.AskOne(prompt, &input)
	color.Unset()

	switch input {
	case 0:
		rifle := afCounterTerrorist.createRifle()
		player2.SetWeaponStrategy(rifle)
	case 1:
		pistol := afCounterTerrorist.createPistol()
		player2.SetWeaponStrategy(pistol)
	case 2:
		knife := afCounterTerrorist.createKnife()
		player2.SetWeaponStrategy(knife)
		//default:
		//	continue
	}
	//	break WEAPONCTerrorst
	//}

	mess, ok := player1.Kill() // Attempt to kill opponent
	gameInstance.notifyAll(fmt.Sprintf("%s %s: %s :%s %s\n", player1.GetTeam(), player1.GetName(), mess, player2.GetTeam(), player2.GetName()))
	fmt.Printf("%s %s: %s :%s %s\n", player1.GetTeam(), player1.GetName(), mess, player2.GetTeam(), player2.GetName())
	if ok {
		fmt.Printf("Terrorist WON!!!\n  %s", mess)
		return
	}
	messi, oki := player2.Kill() // Attempt to kill opponent
	gameInstance.notifyAll(fmt.Sprintf("%s %s: %s :%s %s\n", player2.GetTeam(), player2.GetName(), messi, player1.GetTeam(), player1.GetName()))
	fmt.Printf("%s %s: %s :%s %s\n", player2.GetTeam(), player2.GetName(), messi, player1.GetTeam(), player1.GetName())
	if oki {
		fmt.Printf("Counter Terrorist WON!!!\n  %s", messi)
		return
	}

	player1.SetEquipmentStrategy(afTerrorist.createEquipment())
	player2.SetEquipmentStrategy(afCounterTerrorist.createEquipment())

	player1.UseEquipment()
	player2.UseEquipment()
}
