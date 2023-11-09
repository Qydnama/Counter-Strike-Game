package main

import (
	"fmt"
	"strconv"
)

type DefuseKitStrategy struct {
	Bomb  *BombStrategy
	Exist bool
}

func (e *DefuseKitStrategy) UseEquipment() {
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
				fmt.Println("BombStrategy was defused\nCounter-Terrorist WON!!!")
				return
			} else {
				fmt.Println("BombStrategy was bababahbhashdhsaklfhio asncirwnurioasnfo jsakoaj\nTerrorist WON!!!")
				return
			}
		} else if e.Bomb.babah {
			fmt.Println("BombStrategy was bababahbhashdhsaklfhio asncirwnurioasnfo jsakoaj YOu lose the game\nTerrorist WON!!!")
			return
		}
	}

}
