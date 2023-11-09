package main

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
)

type DefuseKitStrategy struct {
	Bomb *Bomb
}

func (e *DefuseKitStrategy) UseEquipment() {
	color.Set(color.FgBlue)
	fmt.Println("Using the defuse kit")

	bombCode := e.Bomb.Code
	var code string

	for {
		fmt.Printf("Input code:\t %v\n", bombCode)
		fmt.Scan(&code)
		codeInt, err := strconv.Atoi(code)
		if err != nil {
			continue
		}
		if codeInt == bombCode {
			if !e.Bomb.babah {
				e.Bomb.timer.Stop()
				fmt.Println("Counter-Terrorist WON!!!\n  Bomb was defused")
				color.Unset()
				return
			} else {
				color.Unset()
				return
			}
		} else if e.Bomb.babah {
			color.Unset()
			return
		}
	}

}
