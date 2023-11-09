package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

type KnifeCounterTerrorist struct {
}

func (k *KnifeCounterTerrorist) Kill() (string, bool) {
	//fmt.Println("Using a knife as a weapon")
	rand.Seed(time.Now().UnixNano())
	randomNumber1 := rand.Intn(7) + 1
	color.Set(color.FgBlue)
	fmt.Print("Write number between 1 and 7 (shooting to the Terrorist): ")
	var side int
LOOP:
	for {
		fmt.Scan(&side)
		if side >= 1 && side <= 7 {
			break LOOP
		}
		fmt.Print("Number must be between 1 and 7: ")
	}
	color.Unset()
	if randomNumber1 == side {
		return "knifed", true
	} else {
		return "missed with a Knife", false
	}
}
