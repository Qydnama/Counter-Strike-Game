package main

import "fmt"

// concrete observer
type User struct {
	id string
}

func (c *User) update(player string) {
	fmt.Printf("Sending information to User %s about player: %s\n", c.id, player)
}

func (c *User) getID() string {
	return c.id
}
