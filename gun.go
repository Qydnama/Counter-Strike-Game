package main

import (
	"fmt"
)

type GunStrategy struct{}

func (g *GunStrategy) UseWeapon() {
	fmt.Println("Using a gun as a weapon")
}
