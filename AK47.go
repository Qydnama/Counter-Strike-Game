package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

type AK47 struct {
}

func (a *AK47) Kill() (string, bool) {
	//fmt.Println("AK47 shot: PAh PAh")
	rand.Seed(time.Now().UnixNano())
	color.Set(color.FgHiRed)
	randomNumber1 := rand.Intn(3) + 1
	fmt.Print("Write number between 1 and 3 (shooting to the Counter-Terrorist): ")
	var side int
LOOP:
	for {
		fmt.Scan(&side)
		if side >= 1 && side <= 3 {
			break LOOP
		}
		fmt.Print("Number must be between 1 and 3: ")
	}
	color.Unset()
	if randomNumber1 == side {
		return "killed with AK47", true
	} else {
		return "missed a bullet by the AK47", false
	}
}
