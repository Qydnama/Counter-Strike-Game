package main

func GetTerrorist() (IPlayer, error) {
	return newTerrorist(), nil
}

func GetCounterTerrorist() (IPlayer, error) {
	return newCounterTerrorist(), nil
}
