package main

import "fmt"

func GetPlayer(team string) (IPlayer, error) {
	if team == "terrorist" {
		return newTerrorist(), nil
	}
	if team == "counter-terrorist" {
		return newCounterTerrorist(), nil
	}
	return nil, fmt.Errorf("Wrong Team")
}
