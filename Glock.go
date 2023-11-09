package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

type Glock struct {
}

func (a *Glock) Kill() (string, bool) {
	rand.Seed(time.Now().UnixNano())
	randomNumber1 := rand.Intn(5) + 1
	color.Set(color.FgRed)
	var side int
	fmt.Print("Write number between 1 and 5 (shooting to the Counter-Terrorist): ")
LOOP:
	for {
		fmt.Scan(&side)
		if side >= 1 && side <= 5 {
			break LOOP
		}
		fmt.Print("Number must be between 1 and 5: ")
	}
	color.Unset()
	if randomNumber1 == side {
		return "killed with Glock", true
	} else {
		return "missed a bullet by the Glock", false
	}
}
