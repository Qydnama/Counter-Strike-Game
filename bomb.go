package main

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"strings"
	"time"
)

type Bomb struct {
	babah   bool
	Code    int
	Seconds time.Duration
	timer   *time.Timer
}

func (b *Bomb) UseEquipment() {
	color.Set(color.FgRed)
	fmt.Println("Planting the bomb")
	b.Seconds = 5
	var code string
	fmt.Print("Input code for bomb (10 digit): ")
LOOP1:
	for {
		fmt.Scan(&code)
		if len(code) != 10 {
			fmt.Print("Code must be 10 digit number: ")
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

	go func(b *Bomb) {
		for {
			select {
			case _ = <-timer.C:
				b.babah = true
				color.Set(color.FgRed)
				fmt.Println("\nBomb was exploaded(vzarvalos)(babah koroche).\nTerrorist WON!!!")
				color.Unset()
			}
		}
	}(b)
	color.Unset()
}
