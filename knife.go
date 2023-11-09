package main

import (
	"fmt"
)

type KnifeStrategy struct{}

func (k *KnifeStrategy) UseWeapon() {
	fmt.Println("Using a knife as a weapon")
}
