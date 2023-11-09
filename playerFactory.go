package main

import "fmt"

func GetPlayer(team, name string) (IPlayer, error) {
	if team == "terrorist" {
		return newTerrorist(name), nil
	}
	if team == "counter-terrorist" {
		return newCounterTerrorist(name), nil
	}
	return nil, fmt.Errorf("Wrong Team")
}
