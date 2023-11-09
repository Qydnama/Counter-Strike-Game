package main

type Terrorist struct {
	Player
}

func newTerrorist(name string) IPlayer {
	temp := &Terrorist{
		Player: Player{name: name, team: "Terrorist"}}
	gameInstance.Players = append(gameInstance.Players, temp)
	return temp
}
