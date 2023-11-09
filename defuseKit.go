package main

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
)

type DefuseKitStrategy struct {
	Bomb  *Bomb
	Exist bool
}

func (e *DefuseKitStrategy) UseEquipment() {
	color.Set(color.FgBlue)
	if e.Exist {
		fmt.Println("Using the defuse kit")
	} else {
		fmt.Println("Without defuse kit")

	}
	bombCode := e.Bomb.Code
	fmt.Println("BombCode", bombCode)
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
