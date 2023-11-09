package main

import (
	"fmt"
	"sync"
)

type GameInstance struct {
	Players []IPlayer
}

var mu = &sync.Mutex{}

var instance *GameInstance

func GetGameInstance() *GameInstance {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			fmt.Println("Singleton: Creating Players now.")
			instance = &GameInstance{}
		} else {
			fmt.Println("Players already created.")
		}
	} else {
		fmt.Println("Players already created.")
	}

	return instance
}
