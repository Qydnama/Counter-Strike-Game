package main

type CounterTerrorist struct {
	Player
}

func newCounterTerrorist() IPlayer {
	temp := &CounterTerrorist{
		Player: Player{team: "Counter-Terrorist"}}
	gameInstance.Players = append(gameInstance.Players, temp)
	return temp
}
