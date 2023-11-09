package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

type KnifeTerrorist struct {
}

func (k *KnifeTerrorist) Kill() (string, bool) {
	//fmt.Println("Using a knife as a weapon")
	rand.Seed(time.Now().UnixNano())
	randomNumber1 := rand.Intn(7) + 1
	color.Set(color.FgRed)
	fmt.Print("Write number between 1 and 7 (to shoot the Counter-Terrorist): ")
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
