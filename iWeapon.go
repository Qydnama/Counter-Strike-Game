package main

type IWeapon interface {
	Kill() (string, bool)
}
