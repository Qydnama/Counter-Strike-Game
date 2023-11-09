package main

import "fmt"

// concrete observer
type User struct {
	id    string
	brain []string
}

func (c *User) update(comment string) {
	c.brain = append(c.brain, comment)
}

func (c *User) getID() string {
	return c.id
}

func (c *User) getBrainInformation() {
	for _, m := range c.brain {
		fmt.Println(m)
	}
}
