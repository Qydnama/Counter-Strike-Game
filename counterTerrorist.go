package main

type CounterTerrorist struct {
	Player
}

func newCounterTerrorist(name string) IPlayer {
	temp := &CounterTerrorist{
		Player: Player{name: name, team: "Counter-Terrorist"}}
	gameInstance.Players = append(gameInstance.Players, temp)
	return temp
}
