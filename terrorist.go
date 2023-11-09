package main

type Terrorist struct {
	Player
}

func newTerrorist() IPlayer {
	temp := &Terrorist{
		Player: Player{team: "Terrorist"}}
	gameInstance.Players = append(gameInstance.Players, temp)
	return temp
}
