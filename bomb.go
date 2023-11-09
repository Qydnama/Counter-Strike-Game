package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type BombStrategy struct {
	babah   bool
	Code    int
	Seconds time.Duration
	timer   *time.Timer
}

func (b *BombStrategy) UseEquipment() {
	fmt.Println("Planting the bomb")

	var code string
LOOP1:
	for {
		fmt.Print("Input code for bomb (10 digit): ")
		fmt.Scan(&code)
		if len(code) != 10 {
			fmt.Println("Code must be 10 digit number.")
			continue
		}
		codeInt, err := strconv.Atoi(strings.Trim(code, " "))
		if err != nil {
			continue
		}
		b.Code = codeInt
		break LOOP1
	}
	timer := time.NewTimer(b.Seconds * time.Second)
	b.timer = timer

	go func(b *BombStrategy) {
		for {
			select {
			case _ = <-timer.C:
				b.babah = true
				fmt.Println("BombStrategy was exploaded(vzarvalos)(babah koroche).\nTerrorist WON!!!")
			}
		}
	}(b)

}
