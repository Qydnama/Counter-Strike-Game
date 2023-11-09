package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

type M4 struct {
}

func (a *M4) Kill() (string, bool) {
	color.Set(color.FgBlue)
	rand.Seed(time.Now().UnixNano())
	randomNumber1 := rand.Intn(3) + 1
	var side int
	fmt.Print("Input number between 1 and 3(to shoot the Terrorist): ")
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
		return "killed with M4", true
	} else {
		return "missed a bullet by the M4", false
	}
}
